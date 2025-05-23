package aws

import (
	"bytes"
	"fmt"
	"time"

	"github.com/nuzur/nem/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

type AWSClient struct {
	session *session.Session
	config  *config.AWS
}

func New(config *config.AWS) (*AWSClient, error) {
	awsConfig := config
	if awsConfig == nil {
		awsConfigProvider, err := getAwsConfigFromProvider()
		if err != nil {
			return nil, err
		}
		awsConfig = awsConfigProvider
	}

	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(awsConfig.Region),
			Credentials: credentials.NewStaticCredentials(
				awsConfig.KeyID,
				awsConfig.Secret,
				"", // a token will be created when the session it's used.
			),
		},
	)
	if err != nil {
		return nil, fmt.Errorf("error creating aws sesion: %v", err.Error())
	}

	return &AWSClient{
		session: sess,
		config:  awsConfig,
	}, nil
}

func (c *AWSClient) Config() *config.AWS {
	return c.config
}

func (c *AWSClient) UploadToS3(filename string, data []byte) (*s3manager.UploadOutput, error) {
	uploader := s3manager.NewUploader(c.session)

	bucket := c.config.Bucket

	//upload to the s3 bucket
	up, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Body:   bytes.NewReader(data),
	})
	if err != nil {
		return nil, fmt.Errorf("error uploading image to s3: %v %v", err.Error(), up)
	}

	return up, nil
}

func (c *AWSClient) GetSignedS3URL(filename string, expire time.Duration) (*string, error) {
	svc := s3.New(c.session)
	bucket := c.config.Bucket
	_, err := svc.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		return nil, err
	}
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	urlStr, err := req.Presign(expire)
	if err != nil {
		return nil, err
	}

	return &urlStr, nil
}

func (c *AWSClient) FetchFromS3(filename string) ([]byte, error) {

	bucket := c.config.Bucket

	downloader := s3manager.NewDownloader(c.session)
	buf := aws.NewWriteAtBuffer([]byte{})
	_, err := downloader.Download(buf, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c *AWSClient) FetchSecret(secretName string) (*secretsmanager.GetSecretValueOutput, error) {
	svc := secretsmanager.New(c.session)

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}

	return svc.GetSecretValue(input)
}

func (c *AWSClient) CreateSecret(secretName string, secretContent string, KmsKeyId *string) (*secretsmanager.CreateSecretOutput, error) {
	svc := secretsmanager.New(c.session)

	input := &secretsmanager.CreateSecretInput{
		Name:         &secretName,
		KmsKeyId:     KmsKeyId,
		SecretString: &secretContent,
	}

	return svc.CreateSecret(input)
}

func (c *AWSClient) UpdateSecret(secretName string, secretContent string, KmsKeyId *string) (*secretsmanager.UpdateSecretOutput, error) {
	svc := secretsmanager.New(c.session)

	input := &secretsmanager.UpdateSecretInput{
		SecretId:     &secretName,
		KmsKeyId:     KmsKeyId,
		SecretString: &secretContent,
	}

	return svc.UpdateSecret(input)
}

func getAwsConfigFromProvider() (*config.AWS, error) {
	provider, err := config.New()
	if err != nil {
		return nil, fmt.Errorf("error loading config: %v", err.Error())
	}
	var awsConfigs []config.AWS
	if err = provider.Get("aws").Populate(&awsConfigs); err != nil {
		return nil, fmt.Errorf("error reading aws config: %v", err.Error())
	}
	if len(awsConfigs) == 0 {
		return nil, fmt.Errorf("aws config not found")
	}
	awsConfig := awsConfigs[0]

	return &awsConfig, nil
}
