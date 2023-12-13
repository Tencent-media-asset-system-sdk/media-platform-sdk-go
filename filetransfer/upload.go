package filetransfer

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sync"
	"time"

	"github.com/panjf2000/ants"
)

// 上传分辨大小 8M
const BloackSzie = 8 * 1024 * 1024

type retryFunc func() error

func retry(f retryFunc) (err error) {
	sleepTime := 100 * time.Millisecond
	for i := 0; i < 3; i++ {
		if err = f(); err == nil {
			return nil
		}
		time.Sleep(sleepTime)
		sleepTime *= 2
	}
	return err
}

// BatchUploadPart 分片上传文件
func (c Client) batchUploadPart(ctx context.Context,
	filePath, bucket, key, uploadID string, coroutineNum int) (err error) {
	defer ants.Release()
	wg := sync.WaitGroup{}
	// 线程池
	pool, _ := ants.NewPoolWithFunc(coroutineNum, func(v interface{}) {
		var e error
		filebuf := v.([]interface{})[0].([]byte)
		partNumber := v.([]interface{})[1].(int)
		maxTry := 5
		for i := 0; i < maxTry; i++ {
			_, e = c.UploadPart(ctx, bucket, key, uploadID, partNumber, filebuf)
			if e == nil {
				break
			}
		}
		if e != nil && err == nil {
			err = e
		}
		wg.Done()
	})
	defer pool.Release()
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	buffer := make([]byte, BloackSzie)
	// Submit uploadpart one by one.
	partNumber := 1
	for {
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		wg.Add(1)
		filebuf := make([]byte, n)
		copy(filebuf, buffer)
		_ = pool.Invoke([]interface{}{filebuf, partNumber})
		partNumber += 1
	}
	wg.Wait()
	if err != nil {
		return err
	}
	return nil
}

func (c Client) uploadBigFile(ctx context.Context, localPath, key string) (
	uploadId, bucket string, err error) {
	// 上传大文件到数据中心
	bucket = ""
	// 1. 基于 key 创建 uploadId
	uploadId, err = c.CreateMultipartUpload(ctx, bucket, key)
	if err != nil {
		return uploadId, bucket, fmt.Errorf("uploadBigFile failed, %v", err)
	}
	// 2. 批量上传分片
	if err = c.batchUploadPart(ctx, localPath, bucket, key, uploadId, 4); err != nil {
		return uploadId, bucket, fmt.Errorf("uploadBigFile failed, %v", err)
	}
	// 3. 合并文件
	if err = c.CompleteMultipartUpload(ctx, bucket, key, uploadId); err != nil {
		return uploadId, bucket, fmt.Errorf("uploadBigFile failed, %v", err)
	}
	return uploadId, bucket, nil
}

// UploadFile 上传文件
func (c Client) UploadFile(ctx context.Context, localPath, fileName string) (
	key, uploadId, bucket string, err error) {
	file, err := os.Open(localPath)
	if err != nil {
		return key, uploadId, bucket, err
	}
	defer file.Close()
	stat, err := os.Stat(localPath)
	if err != nil {
		return key, uploadId, bucket, err
	}
	defer func() {
		if err != nil {
			c.CreateDeleteFileTask(ctx, []string{key}, true)
			err = fmt.Errorf("UploadFile path: %s error: %v", localPath, err)
		}
	}()
	retry(
		func() error {
			key, err = c.DescribeUploadKey(ctx, fileName, uint64(stat.Size()))
			return err
		})
	fmt.Print("key = ", key)
	if err != nil {
		return key, uploadId, bucket, err
	}
	if stat.Size() > BloackSzie {
		// 大文件分片上传
		uploadId, bucket, err = c.uploadBigFile(ctx, localPath, key)
		return key, uploadId, bucket, err
	} else {
		// 小文件直接 putObject
		filebuf, err := ioutil.ReadAll(file)
		if err != nil {
			return key, uploadId, bucket, err
		}
		retry(
			func() error {
				_, err = c.PutObject(ctx, key, filebuf)
				return err
			})
		return key, uploadId, bucket, err
	}
}

// UploadBuf 上传内存文件
func (c Client) UploadBuf(ctx context.Context, buf []byte, fileName string) (
	key string, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("UploadBuf buf error: %v", err)
		}
	}()
	retry(
		func() error {
			key, err = c.DescribeUploadKey(ctx, fileName, uint64(len(buf)))
			return err
		})
	if err != nil {
		return key, err
	}
	retry(
		func() error {
			_, err = c.PutObject(ctx, key, buf)
			return err
		})
	if err != nil {
		return key, err
	}
	return key, nil
}
