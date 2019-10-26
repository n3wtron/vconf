package vconf

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/n3wtron/vconf/v2/mocks"
	"github.com/phayes/freeport"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"testing"
)

func TestMain(m *testing.M) {
	logrus.SetLevel(logrus.DebugLevel)
	m.Run()
}

func TestServer_GetFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	appName := "app1"
	appVersion := "1.0.0"
	fileName := "config.yaml"
	fileContent := []byte(uuid.New().String())

	mockRepo := mocks.NewMockRepo(ctrl)
	mockRepo.EXPECT().GetFile(appName, appVersion, fileName).Return(fileContent, nil)

	srv := NewServer(mockRepo)

	freePort, err := freeport.GetFreePort()
	assert.NoError(t, err)
	address := fmt.Sprintf("127.0.0.1:%d", freePort)

	go func() {
		err := srv.Start(address)
		if err != nil {
			assert.FailNow(t, "error starting grpc server %s", err)
		}
	}()
	defer srv.Stop()
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	assert.NoError(t, err)
	cl := NewConfigurationClient(conn)
	req := &GetFileRequest{
		AppName:    appName,
		AppVersion: appVersion,
		FilePath:   fileName,
	}
	resp, err := cl.GetFile(context.TODO(), req)
	assert.NoError(t, err)
	assert.Equal(t, resp.FileContent, fileContent)
}