package sqs

import (
	"strconv"

	"list_product/config"
	awsService "list_product/services/aws"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

//Messager --
type Messager struct {
	SQSClient sqsiface.SQSAPI
}

//New --
func New() *Messager {
	client := sqs.New(awsService.GetAwsSession(config.ENV.AwsRegion))
	return &Messager{
		SQSClient: client,
	}
}

func getQueueURL(service sqsiface.SQSAPI, queueName string) (*sqs.GetQueueUrlOutput, error) {
	return service.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(queueName),
	})
}

//SendMessage --
func (m *Messager) SendMessage(body string) error {
	qURL, err := getQueueURL(m.SQSClient, config.ENV.SQSQueueURL)
	if err != nil {
		return err
	}

	_, err = m.SQSClient.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageBody:  aws.String(body),
		QueueUrl:     qURL.QueueUrl,
	})

	if err != nil {
		return err
	}

	return nil
}

//ReadMessages -- Pool messagem for given queue
func (m *Messager) ReadMessages(queueName string, maxMessages int64) ([]*sqs.Message, error) {
	queueURL, err := getQueueURL(m.SQSClient, queueName)
	if err != nil {
		return nil, err
	}

	result, err := m.receiveMessagesFromQueue(queueURL, maxMessages)
	if err != nil {
		return nil, err
	}

	return result.Messages, nil
}

func (m *Messager) receiveMessagesFromQueue(qURL *sqs.GetQueueUrlOutput, maxMessages int64) (*sqs.ReceiveMessageOutput, error) {
	sleepTime, err := strconv.ParseInt(config.ENV.SQSQueueSleepTime, 10, 64)
	if err != nil {
		panic(err)
	}

	return m.SQSClient.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:            qURL.QueueUrl,
		MaxNumberOfMessages: aws.Int64(maxMessages),
		WaitTimeSeconds:     aws.Int64(sleepTime),
	})
}

//DeleteMessages --
func (m *Messager) DeleteMessages(messages []*sqs.Message, queueName string) error {

	for _, message := range messages {
		err := m.DeleteMessage(message, queueName)

		if err != nil {
			return err
		}
	}

	return nil
}

//DeleteMessage --
func (m *Messager) DeleteMessage(message *sqs.Message, queueName string) error {
	qURL, err := getQueueURL(m.SQSClient, queueName)
	if err != nil {
		return err
	}

	_, errDel := m.SQSClient.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      qURL.QueueUrl,
		ReceiptHandle: message.ReceiptHandle,
	})
	return errDel
}
