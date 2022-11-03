package data

import "errors"

// DataTypeEnum
type DataTypeEnum string

const (
	// DataTypeS3
	DataTypeS3 DataTypeEnum = "s3"
	// DataTypeSFTP
	DataTypeSFTP DataTypeEnum = "sftp"
	// DataTypeLocal
	DataTypeLocal DataTypeEnum = "local"
)

func (o DataTypeEnum) IsValid() bool {
	return (o == DataTypeS3 ||
		o == DataTypeSFTP ||
		o == DataTypeLocal)
}

const (
	// Errors
	ERR_INVALID_DATA_KEY     = "invalid data key"
	ERR_UNEXPECTED           = "unexpected error"
	ERR_DATA_MOVER_NOT_FOUND = "Data mover not found"

	// Data results
	S3_DATA_RESULT    = "s3 data"
	SFTP_DATA_RESULT  = "sftp data"
	LOCAL_DATA_RESULT = "local data"

	// Data save success messages
	S3_SAVE_SUCCESS    = "succesfully saved in s3"
	SFTP_SAVE_SUCCESS  = "successfully saved in sftp"
	LOCAL_SAVE_SUCCESS = "successfully saved locally"
)

/** Database **/

type database struct{}

func NewDatabase() *database {
	return &database{}
}

func (d *database) RetrieveData(dataKey DataTypeEnum) (string, error) {
	if !dataKey.IsValid() {
		return "", errors.New(ERR_INVALID_DATA_KEY)
	}

	switch dataKey {
	case DataTypeS3:
		return S3_DATA_RESULT, nil
	case DataTypeSFTP:
		return SFTP_DATA_RESULT, nil
	case DataTypeLocal:
		return LOCAL_DATA_RESULT, nil
	default:
		return "", errors.New(ERR_UNEXPECTED)
	}
}

/** Data Service **/

type DataService struct {
	db        *database
	dataMover DataMover
	dataKey   DataTypeEnum
}

func NewDataService(db *database, dataKey DataTypeEnum) *DataService {
	return &DataService{
		db:        db,
		dataMover: dataMovers[dataKey],
		dataKey:   dataKey,
	}
}

func (dm *DataService) Save() (string, error) {
	result, err := dm.db.RetrieveData(dm.dataKey)
	if err != nil {
		return "", err
	}
	return dm.dataMover.MoveData(result)
}

/** Data movers **/

type DataMover interface {
	MoveData(data string) (string, error)
}

var dataMovers = map[DataTypeEnum]DataMover{
	DataTypeS3:    &s3Uploader{},
	DataTypeSFTP:  &sftpUploader{},
	DataTypeLocal: &localStore{},
}

type s3Uploader struct{}

func (s *s3Uploader) MoveData(data string) (string, error) {
	return S3_SAVE_SUCCESS, nil
}

type sftpUploader struct{}

func (s *sftpUploader) MoveData(data string) (string, error) {
	return SFTP_SAVE_SUCCESS, nil
}

type localStore struct{}

func (s *localStore) MoveData(data string) (string, error) {
	return LOCAL_SAVE_SUCCESS, nil
}
