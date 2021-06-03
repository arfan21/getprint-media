package dropbox

import (
	"crypto/rand"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"path/filepath"

	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/sharing"
)

type DropboxSdkFiles interface {
	Upload(arg *files.CommitInfo, content io.Reader) (res *files.FileMetadata, err error)
	DeleteV2(arg *files.DeleteArg) (res *files.DeleteResult, err error)
}

type DropboxSdkSharing interface {
	CreateSharedLinkWithSettings(arg *sharing.CreateSharedLinkWithSettingsArg) (res sharing.IsSharedLinkMetadata, err error)
}

type Dropbox interface {
	Uploader(file *multipart.FileHeader) (path string, err error)
	CreateSharedLink(path string) (string, error)
	Delete(path string) error
}

type dbx struct {
	token   string
	files   DropboxSdkFiles
	sharing DropboxSdkSharing
}

func NewDropboxRepository() Dropbox {
	token := os.Getenv("DROPBOX_ACCESS_TOKEN")
	dbxConfig := dropbox.Config{
		Token: token,
	}
	dbxFiles := files.New(dbxConfig)
	dbxSharing := sharing.New(dbxConfig)

	return &dbx{token, dbxFiles, dbxSharing}
}

func (repo dbx) Uploader(file *multipart.FileHeader) (path string, err error) {
	oldFilename := file.Filename
	extensionFile := filepath.Ext(oldFilename)
	randArrayByte := make([]byte, 10)
	if _, err := rand.Read(randArrayByte); err != nil {
		return "", err
	}

	randString := fmt.Sprintf("%X", randArrayByte)
	filename := randString + extensionFile

	srcFile, err := file.Open()
	if err != nil {
		return "", err
	}
	defer srcFile.Close()

	args := &files.CommitInfo{
		Path: "/getprint/" + filename,
		Mode: &files.WriteMode{
			Tagged: dropbox.Tagged{Tag: "add"},
		},
		Autorename:     true,
		Mute:           false,
		StrictConflict: false,
	}
	res, err := repo.files.Upload(args, srcFile)
	if err != nil {
		return "", err
	}

	return res.PathLower, nil
}

func (repo dbx) CreateSharedLink(path string) (string, error) {
	args := &sharing.CreateSharedLinkWithSettingsArg{
		Path:     path,
		Settings: nil,
	}

	res, err := repo.sharing.CreateSharedLinkWithSettings(args)
	if err != nil {
		return "", err
	}
	linkRes := res.(*sharing.FileLinkMetadata)

	parsedUrl, err := url.Parse(linkRes.Url)
	if err != nil {
		return "", err
	}

	fixURL := fmt.Sprintf("%s://%s%s?raw=1", parsedUrl.Scheme, parsedUrl.Host, parsedUrl.Path)

	return fixURL, nil
}

func (repo dbx) Delete(path string) error {
	args := &files.DeleteArg{
		Path: path,
	}
	_, err := repo.files.DeleteV2(args)
	if err != nil {
		return err
	}

	return nil
}
