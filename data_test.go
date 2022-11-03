package data_test

import (
	"data"
	"testing"
)

func TestRetrieveData(t *testing.T) {
	db := data.NewDatabase()

	// Bad input
	result, err := db.RetrieveData("bad-input")
	if err == nil {
		t.Errorf("Expected err %s, got nil", data.ERR_INVALID_DATA_KEY)
	}
	if result != "" {
		t.Errorf("Result should be empty on bad input, got %s", result)
	}

	// Valid inputs
	tcs := []struct {
		dataKey data.DataTypeEnum
		want    string
	}{
		{
			dataKey: data.DataTypeS3,
			want:    data.S3_DATA_RESULT,
		},
		{
			dataKey: data.DataTypeSFTP,
			want:    data.SFTP_DATA_RESULT,
		},
		{
			dataKey: data.DataTypeLocal,
			want:    data.LOCAL_DATA_RESULT,
		},
	}

	for _, tc := range tcs {
		got, err := db.RetrieveData(tc.dataKey)
		if err != nil {
			t.Fatal(err)
		}

		if got != tc.want {
			t.Errorf("got %s, want %s", got, tc.want)
		}
	}
}

func TestSaveData(t *testing.T) {
	db := data.NewDatabase()

	tcs := []struct {
		dataService *data.DataService
		want        string
	}{
		{
			dataService: data.NewDataService(db, data.DataTypeS3),
			want:        data.S3_SAVE_SUCCESS,
		},
		{
			dataService: data.NewDataService(db, data.DataTypeSFTP),
			want:        data.SFTP_SAVE_SUCCESS,
		},
		{
			dataService: data.NewDataService(db, data.DataTypeLocal),
			want:        data.LOCAL_SAVE_SUCCESS,
		},
	}

	for _, tc := range tcs {
		got, err := tc.dataService.Save()
		if err != nil {
			t.Fatal(err)
		}

		if got != tc.want {
			t.Errorf("got %s, want %s", got, tc.want)
		}
	}
}
