package cli_test

import (
	"fmt"
	"testing"

	"github.com/filipe/exagonal/adapters/cli"
	mock_application "github.com/filipe/exagonal/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Product 1"
	productPrice := 10.0
	productStatus := "enabled"
	productID := "abc"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productID).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productID).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s\n", productID, productName, productPrice, productStatus)

	result, err := cli.Run(service, "create", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	result, err = cli.Run(service, "get", productID, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	result, err = cli.Run(service, "enable", productID, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	result, err = cli.Run(service, "disable", productID, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	result, err = cli.Run(service, "invalid", "", "", 0)
	require.Equal(t, "Invalid action", result)
}
