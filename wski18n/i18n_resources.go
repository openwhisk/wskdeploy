// Code generated by go-bindata.
// sources:
// wski18n/resources/de_DE.all.json
// wski18n/resources/en_US.all.json
// wski18n/resources/es_ES.all.json
// wski18n/resources/fr_FR.all.json
// wski18n/resources/it_IT.all.json
// wski18n/resources/ja_JA.all.json
// wski18n/resources/ko_KR.all.json
// wski18n/resources/pt_BR.all.json
// wski18n/resources/zh_Hans.all.json
// wski18n/resources/zh_Hant.all.json
// DO NOT EDIT!

package wski18n

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _wski18nResourcesDe_deAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesDe_deAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_wski18nResourcesDe_deAllJson,
		"wski18n/resources/de_DE.all.json",
	)
}

func wski18nResourcesDe_deAllJson() (*asset, error) {
	bytes, err := wski18nResourcesDe_deAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wski18n/resources/de_DE.all.json", size: 0, mode: os.FileMode(436), modTime: time.Unix(1536397114, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wski18nResourcesEn_usAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x3c\x6b\x8f\x1b\x37\x92\xdf\xfd\x2b\x0a\xc1\x02\x4e\x80\x19\x4d\xb2\x87\x03\x0e\xc6\xf9\x80\x39\x7b\x92\x9d\x8d\xed\x31\xc6\xe3\x0d\x02\xdb\x68\x53\xdd\x25\x89\xab\x6e\xb2\x97\x64\x4b\x56\x06\xfa\xef\x87\x2a\x92\xdd\xad\x47\x3f\x24\x3b\xb8\xf5\x17\x4b\x22\x59\x2f\x16\xeb\x49\xce\x87\x27\x00\x8f\x4f\x00\x00\xbe\x93\xd9\x77\xcf\xe0\xbb\xc2\xce\x93\xd2\xe0\x4c\x7e\x49\xd0\x18\x6d\xbe\xbb\xf0\xa3\xce\x08\x65\x73\xe1\xa4\x56\x34\xed\x86\xc7\x9e\x00\x6c\x2f\x7a\x20\x48\x35\xd3\x1d\x00\x6e\x69\x68\x68\xbd\xad\xd2\x14\xad\xed\x00\xf1\x2e\x8c\x0e\x41\x59\x0b\xa3\xa4\x9a\x77\x40\xf9\x2d\x8c\x76\x42\x49\x8b\x2c\xc9\xd0\xa6\x49\xae\xd5\x3c\x31\x58\x6a\xe3\x3a\x60\xdd\xf3\xa0\x05\xad\x20\xc3\x32\xd7\x1b\xcc\x00\x95\x93\x4e\xa2\x85\xef\xe5\x04\x27\x17\xf0\x56\xa4\x4b\x31\x47\x7b\x01\xd7\x29\xad\xb3\x17\xf0\x60\xe4\x7c\x8e\xc6\x5e\xc0\x7d\x95\xd3\x08\xba\x74\xf2\x03\x08\x0b\x6b\xcc\x73\xfa\xdf\x60\x8a\xca\xf1\x8a\x15\x63\xb3\x20\x15\xb8\x05\x82\x2d\x31\x95\x33\x89\x19\x28\x51\xa0\x2d\x45\x8a\x93\xd1\xbc\x68\xdd\xc5\xc9\xc3\x02\xe1\xae\x44\xf5\xdb\x42\xda\x25\xbc\x64\x66\x0a\x22\xe1\x41\xeb\xfc\xa3\xfa\xa8\x1e\x34\x4c\x71\x2e\x15\xac\xb5\x59\x4a\x35\x87\xb5\x74\x0b\x58\xdb\xa5\x67\xfc\x02\x4c\xe5\x09\x7c\x5a\xff\xf6\x14\x52\x5d\x14\x42\x65\xcf\x08\xc0\x47\xf7\x97\x66\x3a\x43\x5c\x48\x0b\x6b\x99\xe7\x41\x76\x2d\xfc\xc2\x5a\x74\xb6\xc5\xab\x54\x50\x08\x25\x67\x68\xdd\x64\x23\x8a\x1c\xb4\x69\xfd\x50\xe4\x1f\xd5\xed\x0c\xd2\xca\x18\x22\x39\x93\x06\x53\xa7\xcd\x06\x32\x8d\x56\x39\x58\x88\x15\x82\x50\x9b\x7a\x09\xcc\x64\x8e\x17\x0d\x39\x50\x1a\xa9\x9c\x05\x47\x24\x2d\x30\x2f\xa1\x40\x6b\xc5\x1c\x27\x9e\x50\x84\x42\x5b\xc7\xec\x68\x05\x6b\xb1\xb1\xa0\x67\x50\x59\x96\x43\x0d\xc4\xe9\xc8\x89\x50\xd9\x95\x36\x50\xa9\x2e\xce\x84\x41\x16\xca\x8e\x48\x5a\x5f\xe0\xb2\x80\x52\xb8\xc5\x95\xd3\x57\x3b\x8c\x8f\x9b\x05\x97\x59\x3d\x90\xd5\x7b\x79\x04\x40\xa4\xf0\xf8\xaf\x23\xa9\x18\x9c\xde\x4b\xce\x47\x75\x5d\x29\xb7\xa0\x63\x93\xb2\x3a\x3e\xfb\xa8\x1a\xd8\x06\x45\x66\x21\x35\x98\xd1\x04\x91\x5b\x98\x19\x5d\xc0\x5f\xfe\x76\xf7\xfa\xe6\x6a\xb2\xb6\xcb\xd2\xe8\xd2\xc2\x74\x03\x19\xce\x44\x95\xbb\x8f\xea\x6e\x85\x66\x6d\xa4\xc3\xf8\x13\xa4\x5a\xcd\xe4\x9c\x37\x9d\x8e\xea\x8b\x57\xb7\xcf\x3e\x2a\x80\x1d\x49\x5e\x86\x49\xff\xdd\x9a\xfc\x3f\x3d\x02\xb8\x33\x41\x3d\x37\x20\xf2\x1c\xdc\xc2\x60\x0f\x70\x51\xca\x05\x69\xd0\xdf\xee\xde\x3d\xd0\xd7\xca\x2d\xe0\xd7\x9b\xdf\xe1\xf2\xb2\x3e\xc5\xf0\xe6\xfa\xf5\xcd\xbb\xb7\xd7\x2f\x6e\x3a\xb1\x8e\x38\xe7\x76\xa1\x8d\xeb\x37\x5a\x6f\x8d\x5e\xc9\x0c\x2d\x08\xb0\x55\x51\x08\x43\x52\xa6\xf9\xa4\xd3\x07\x9a\x3a\x45\x52\xf2\x68\xdd\xae\xe2\x5e\x63\x06\x53\x61\x31\x23\x96\x23\x8d\xad\xbd\x85\xdf\xaf\x5f\xbf\x1a\x63\x97\x02\xbd\xdd\x86\xe9\x1a\x9c\xd6\x39\x58\x74\x74\xbe\xf8\x6c\x06\xa9\x6e\x74\x65\x40\x97\xa8\xd6\x4c\x6f\x19\xec\x6c\x38\x96\x62\xf7\xb0\x8f\xa7\x65\x85\xc6\x12\xee\x2e\xe1\x49\xe5\xd8\xce\x85\x79\xa0\xaa\x62\x8a\x86\x64\x57\x6f\xf8\x68\x5c\x76\xa3\xd2\x7e\xbe\x9d\x06\x9a\xe4\x99\x6d\x36\xa7\x66\x76\x8a\x6e\x8d\xa8\x20\xcd\x25\x89\x5d\xa8\x0c\x2c\x9a\x15\x9a\xd1\x4e\x61\x3c\x0d\xad\xed\x25\x3c\x51\x15\xf8\x87\x1d\xd5\xe9\xde\x0a\x5a\xa7\x4b\x82\x2f\xf2\x36\x3c\xda\xa2\x38\x9d\x55\x87\xec\xc2\x4b\x39\x9b\x21\x5b\xf4\x68\x71\x4d\xa5\xc8\x77\x33\x39\xcf\x76\x8d\x10\xfd\x74\xf8\xcb\x48\x0b\xd6\x3b\xb5\x6d\xbd\xce\x87\x71\x59\x1a\xfd\x4f\x4c\x1d\x9d\x77\x78\x7b\x7f\xf7\xf7\x9b\x17\x0f\xa3\xf5\x24\x8a\xba\x63\x9f\xde\x77\xfa\x19\x36\x96\x5e\x21\xc6\xea\xc3\x58\x5c\x06\x0b\xbd\x42\x7b\x88\x73\xbd\x90\xe9\x02\xd6\x68\xb0\x09\x8a\x98\x0e\x3a\x35\x3b\x9a\xb0\x6f\x2f\x76\xe2\x8c\x0c\x73\x74\xb4\xd9\xc7\x99\xda\x01\xe6\xdd\xb9\xa9\xd4\xb3\x7f\x3b\xf7\x76\x1c\xd2\x31\x6d\x80\xef\xb5\xca\x37\x1c\x5f\x59\x98\x69\xd3\x12\x0f\x47\x7f\xac\x60\x85\xce\xf0\x87\xd1\x7a\x83\x5f\x7a\xfc\xc0\x0d\x0f\x42\xa0\x64\x47\xb8\xb5\xc8\xc7\x2a\xcd\x08\x44\x96\xb6\x4b\xcc\x31\xeb\xc7\x48\xd6\x66\x47\x49\x66\x95\xe2\xb8\xd9\xdb\x88\x8e\x78\x8c\x56\x51\x00\xea\xe9\xd8\xd3\x02\xff\x63\x87\xd0\x5b\x9b\xea\xe7\x61\x76\x79\x82\xd3\x9d\xe5\x62\x9e\x88\x52\x26\xe4\xde\x3b\xf8\xf7\xfe\xe9\xfa\xed\x2d\x7c\x26\xff\xff\x79\x24\xc4\x7e\x47\xd4\x02\xfa\x8f\x9b\xfb\x77\xb7\x77\x6f\x46\xc1\xad\xdc\x22\x59\x62\xd7\xe1\xa6\x61\x6d\xe4\x1f\xfc\x03\x7c\xfe\xf5\xe6\xf7\x31\x40\x53\x34\x2e\xa1\xdd\xe9\x80\x4a\xf2\x25\xeb\x4d\x47\x76\x42\x93\x79\x2b\xc7\x00\xe6\x50\xac\x03\x6a\x3b\xa8\xfb\x3e\x46\x7a\xd2\xee\x87\x86\x03\x87\xc5\x4b\x25\xcf\xf5\x3a\x09\x30\xba\xb2\x4f\x9e\x04\xf5\xa4\x61\xa8\xcd\xf1\xed\x93\x4b\x9d\x34\xd4\x7e\x70\x04\xe8\xd2\xe0\x4a\xe2\xba\x03\xae\x5d\x30\xa1\x11\xe8\xd5\x8e\xa3\x2e\x73\xa1\x46\x60\x58\xe2\x66\xf4\x96\x2e\x71\x33\x96\x70\x2f\xe9\x60\x08\x7a\x05\x1d\x8d\x44\x9d\x4e\x3b\x72\x0c\x50\x08\xb3\xc4\x2c\x9a\x92\x51\xa2\x62\x38\x09\x1d\xfa\x2e\x66\x02\x2a\x9e\x32\x0c\x31\x5a\x87\x81\x5d\xdd\x71\x4e\x23\xc0\xd6\x89\x40\x07\xdc\x66\x7c\x34\xd3\x03\x14\xfa\xb8\x20\x47\x6b\xa3\xb4\x47\x80\xb6\xce\xc8\x4e\xc8\x7e\xeb\x2a\x8b\xe4\xbc\x66\x52\x61\x46\x56\xd9\xc9\xa2\x0e\x97\x47\x60\x70\xa6\x5b\x08\x3c\x06\xba\x72\x65\x35\x86\x58\xaf\x6e\x2b\x34\x53\x6d\xbb\x40\x86\xd1\x53\x81\x96\xc2\x88\xa2\x53\xc0\x46\x14\xe8\xd0\xc0\x4a\xe4\x15\xb2\xf7\x26\x63\x0a\xff\xb8\x7e\xf5\xfe\xe6\x33\x39\xf7\x42\x9c\x88\xaa\xef\x34\x7e\xfe\xf9\xf6\xd5\xcd\x67\x4a\x73\x9d\x90\x1c\x20\x1f\xa3\xe0\xef\xef\xee\xde\x0c\xa3\x66\xab\x9a\x14\xd2\x52\x2c\xce\xfe\xa2\xdb\x5d\x90\x23\xa6\x19\x4d\xee\x0e\x64\x0b\xa4\x05\xa5\x63\xd6\x5d\x19\xcc\x26\x1f\xfb\xf6\x7d\x0f\xa3\xcf\x94\x7b\x30\x92\xcf\xe3\x64\xfa\xab\xf0\x0c\x1d\x37\xc2\xd4\xe4\xe6\x67\xa1\x0a\xac\xf4\x55\x45\xf7\xf9\xf9\xf0\xf8\x38\xa1\xcf\xdb\xed\xa7\x0b\x1f\x18\x3d\x3e\x4e\xac\xae\x4c\x8a\xdb\xed\x28\x9c\x7e\xc3\x86\x70\x72\x01\x22\xec\x95\x45\x77\x1e\xae\x5a\x3c\x43\xd8\x76\xe4\x48\x2c\xd6\x3f\x9c\xcf\x67\x29\xe7\xeb\x44\x70\x41\x38\x71\x7a\x89\x6a\x90\x65\x5a\x01\x7e\x05\xf0\x8a\xf3\x98\xaf\x54\x21\x8c\x5d\x88\x3c\xc9\x75\x2a\xf2\xce\x34\x29\xcc\x6a\x45\xb6\xc1\x14\x86\x88\x97\x57\x87\xe3\x39\x12\xa1\x42\x47\xd9\xc1\xd9\x28\xa5\x72\x68\x14\x3a\x10\x8e\xd8\xad\x4c\x3e\xc0\x6b\x13\x37\x24\xa9\x50\x29\xe6\x79\xa7\xd7\xbe\xfb\x75\x02\x2f\xfc\x9c\xa6\x60\xc4\x79\xd0\x48\x04\x33\x21\xbb\xa1\xb7\x0a\xd2\x99\xcc\xc2\x59\x2c\xca\x1c\x1d\x82\xad\x68\x4b\x67\x55\x9e\x6f\x26\x70\x5f\x29\xf8\x7c\x98\x71\x7d\xe6\x04\x81\x33\x56\xb2\x8d\x4e\x8a\x3c\xdf\x34\xe9\xa9\xcf\x44\xc6\x52\xea\xab\x65\x89\x75\xc2\x55\x5d\xd1\xe2\xe5\xe5\xe5\xe5\xf3\xe7\xcf\x9f\x1f\x2f\xaa\xbf\xe3\xa5\x40\x13\x68\xe2\x28\xac\xdc\x1b\xc1\x6c\x8c\x88\xa2\x68\x32\x08\x0d\x15\x2f\x9c\x7e\x25\x3b\x7f\xaf\xdb\x6b\xc7\x23\xe9\xdd\xef\xf7\xed\x90\xb5\x77\xc7\x47\xe3\x1b\x92\xdf\x0e\xca\x33\x24\x18\x7a\x1d\x09\x17\xb1\xd8\x5b\x93\x95\x4b\x84\x4b\x28\xde\xea\x40\xfa\xf8\x38\x49\x8b\x6c\xbb\x0d\xa5\xaf\xc7\xc7\x09\x2d\x74\x9b\x12\xb7\x5b\xb6\x94\xb4\x76\xbb\xfd\x34\x99\xf4\xe2\xe6\x20\x79\x93\x44\x7d\x1e\xe8\xa3\x3d\x3e\x52\xc8\x1e\x10\x10\x91\xdb\xed\x27\x58\x08\x0b\x53\x44\xb5\xc3\x70\x7d\x42\xc6\x63\xef\x6e\xbc\xbd\x8c\xe3\x70\x94\x80\xc9\xa4\xa7\x64\x19\x50\x34\xd5\xe7\x6f\xc7\x62\x03\x73\x0c\x93\x71\x76\x37\x9b\xef\x9b\x19\x47\x19\xed\xe5\x33\xc3\x12\x55\x86\x2a\x3d\x45\x9c\xcd\xa2\xf3\xf1\x34\x47\xa4\x53\xa6\x2f\x8f\xa2\xf9\x1a\xc5\x39\x4e\x05\x19\x86\xca\x74\xc5\x65\x2f\x77\x6a\xce\xc7\x59\xff\x7f\xf4\x11\x91\x9f\xd3\xf4\xe4\xeb\x76\xf0\xd0\xcc\x7d\x9b\x3d\x1c\x79\x32\xba\x28\xe9\xdf\xc7\xf7\x7b\xdd\x83\x73\x76\xb2\x8f\xaa\x50\x21\x38\xd7\xe7\x30\x45\xde\x03\xd4\x15\x88\x3e\x5a\x20\xab\x0c\xed\x64\xac\x71\xb6\x3c\xe2\x9f\xa7\x6f\x91\xc7\x99\xae\x54\x96\x04\x7a\x83\xa5\xea\x54\x80\x50\x55\x3f\x6a\x21\x43\xe9\x9e\x2f\x20\x10\x5d\xad\xc2\x7d\x6c\xae\xef\x17\x71\xd9\x49\xf9\xcf\x04\x41\x58\xe6\x85\xdb\xe3\x63\xc3\x82\x50\x53\x4b\x42\xdb\xa8\xab\xf3\xe6\x47\x39\x99\x80\x56\xbd\xcf\x20\xd7\x31\xb2\x0b\xee\xc3\x36\xe1\x56\xbd\x6d\x44\x87\xa9\x57\x04\x24\x20\x5a\xed\x89\x76\x57\xd3\x5f\x1e\x08\xda\x6f\x7c\xdf\x6d\xe8\xa6\xc5\xcd\xfd\xfd\xdd\xfd\xbb\x0e\xba\x9f\xef\xff\x03\x3f\x1d\x0e\x06\x9e\x3f\xef\x71\x3f\xc6\xec\x1e\xb4\xa5\xd2\x6b\x95\x50\xa4\x30\x7c\xd4\x69\x16\x89\x2a\xac\x9a\x40\xab\x38\xce\x3d\x07\x5b\x95\xbe\x44\x7f\xc5\x65\xe5\x89\xdd\x58\x87\x05\x4c\xa5\xca\xa4\x9a\x5b\xd0\x06\xe6\xd2\x2d\xaa\xe9\x24\xd5\x45\xdd\xde\xeb\xf7\x97\xc6\x44\x9f\x99\x1a\x14\xae\x8b\x4c\xbe\x58\x04\x3c\x65\x47\x2d\xf9\x7a\x09\xdf\x48\x8a\x77\x31\x9e\xd1\x20\x1a\xb3\xdd\x72\x5f\xc0\x8f\xa5\x3a\xf3\x03\xf4\x61\x20\x9b\x69\x91\xe4\xcf\x4a\x2f\x49\xd9\xc1\x49\xf9\x93\x48\x9a\x21\x66\x89\x54\x2b\xbd\xec\x22\xe8\x67\x36\x5b\x64\x2e\xfc\x34\x3e\x90\xb4\x0c\xd6\x0b\x6c\x75\xca\x9c\xbf\x57\x14\x86\xfe\x1c\x6a\x97\xb8\xa9\x6b\x28\x14\xef\x0a\xa7\x4d\x5f\x7d\xa8\x9e\xc3\xe5\x86\x0f\x51\x98\x9f\x48\x1f\x03\x9c\x41\x9c\xb1\x94\x9a\x28\xed\xbc\xb1\xeb\x40\xf8\xba\x5d\x73\x65\x5b\xcd\xb3\x29\xdf\xe5\xa2\x67\x3b\xa2\x1e\x42\xca\xd1\x7b\x21\x6d\x21\x5c\xda\x15\xbe\x13\x83\xb5\x7a\xd0\x82\x8c\x51\x64\xd1\x9e\x4a\xb5\x5f\xdc\xf7\xe3\x81\x06\xbe\x9f\xc4\x64\x32\x12\xde\x56\x36\x6f\x34\xa9\x68\x01\xd9\xa9\x25\xfb\xd1\xc8\x46\x3f\x13\x21\xff\x27\xf5\x12\xb9\xec\x12\xdb\xad\x1f\xe5\x4b\x55\x7e\x4b\xea\xb2\x2d\xe1\x0a\x9f\x89\x96\xa3\x37\xb2\xb8\x59\x49\xb4\x0b\xdf\xa8\xa3\x35\xfe\xe3\x18\x39\x47\x12\x07\x44\x7d\x7f\x0a\x41\x7b\x72\xe5\xa3\xe0\x29\x7a\x6a\xc1\x57\x79\xbc\x28\xf1\x8b\x43\x65\x23\xd1\xf8\x85\x7d\x18\xb1\xf3\x35\xac\xd8\x64\x8e\x5d\x05\xcc\xe6\x28\xcf\xd1\xdf\x23\x09\xb6\xb7\x29\x95\x1f\x74\x44\xc9\xbf\xc9\xb4\x75\x7c\x47\xcb\xd4\x93\x9e\x78\x8e\xf9\xf4\xd4\xd8\x3a\xe8\xdb\x61\x98\xe3\x42\x12\x63\x23\x65\xa1\x36\xb5\x6e\x90\x11\x69\x6d\xfb\xa0\x5c\x43\x11\xb5\x26\x61\x90\x8d\xca\xe4\xa7\x6b\xae\x2f\x6c\x85\x14\xfa\xfd\xfd\x2b\xa6\x80\x4b\x5d\x7c\x94\x3e\xec\xe4\xd8\x9f\xfc\xe5\xa0\x31\x84\x14\x22\x9f\x69\x53\x74\x4a\xee\x75\x1c\xef\xa3\x60\x02\x0f\x66\x03\x62\x2e\xa4\x1a\x4a\xe9\x8d\x49\xfe\x69\xb5\xaa\x8d\x6d\x5a\x64\x3d\x9d\x5b\x2e\xee\x4b\x55\x56\x0e\x32\xe1\x04\xbc\x0e\xd2\x78\x9a\x16\xd9\x53\x32\xbd\xfd\x98\x44\x29\x9b\x0a\xbc\x57\x1a\x6d\x12\x8b\xff\xaa\x50\x75\x96\xc8\xfd\x7d\xd6\xab\x77\x61\xd6\xee\x61\x69\xd9\x77\xaf\xcf\x7b\x97\x35\xae\xdf\xde\xfa\x05\xa5\xa4\xd9\xa9\x50\x3e\x14\x99\xa2\x0f\x06\xda\x17\xcc\x1a\x25\xbb\x8a\x24\x1d\x81\x39\x81\xb7\x39\x0a\x8b\x50\x95\x99\x70\x7b\xb7\x43\xbc\xf3\x4c\xf3\x2a\xdb\xa7\x53\x58\x10\xb0\xc6\xe9\x3e\x86\xc1\xdd\x09\x72\xea\x57\xd0\xeb\x23\x76\x84\x44\x13\x56\x4d\xe0\xd6\xf9\xec\x4b\xbb\x05\xfb\xe2\xdd\x3b\x0f\xf5\xc1\xbb\xf0\xd2\xd1\x0a\x43\xdb\xb5\x20\x28\xf8\xa5\xc4\x74\xcc\x49\x0a\xb4\xc6\x2d\x8e\xf6\x81\x0c\x63\x42\x58\xbf\x92\x7a\x26\xbc\x31\x12\x04\x56\x57\xae\x6d\x2c\x26\xf0\x5b\x63\x84\xa3\xa9\xa0\x65\x17\xb5\x39\x91\xb6\x09\x16\x06\xdc\x5a\x60\x27\x8a\x29\xa1\x6c\xc5\x61\x92\x49\x33\xca\xc8\x1d\x65\x8b\xf8\xa8\xe5\x5e\x6a\xa9\x7c\x48\xe5\x53\x34\x87\xad\x5b\xc5\xcd\x71\xbe\xa0\x14\x30\x72\xc5\xb7\x7a\xf7\x2c\x5c\x3f\x1b\xa9\xa0\x84\x5d\xac\x30\xc9\x74\xba\xc4\xae\xbb\xf7\x2f\x84\x62\xa8\x62\x85\xf0\x92\x27\x82\x2c\x38\x00\x1f\x08\x2c\x65\x8e\x89\xc8\x0d\x8a\x6c\x93\xe0\x17\x69\x3b\xef\x36\xfc\x4c\x27\x24\xcc\x04\x3f\x73\x00\x76\x16\xef\xe6\x35\x59\x89\x44\xeb\x15\xca\x52\xe4\x94\x8b\x29\x76\x35\x47\xee\x14\x02\xe9\x61\x8e\xfb\x69\x7f\xf3\x35\x6e\x89\x5b\x6b\xa8\x91\x71\xd3\xc4\xcb\x9a\x66\xc7\x6f\xde\xb0\x2e\xa4\x85\xa5\x54\x19\x1d\x90\xa0\x8b\xa1\x29\x79\xe0\x78\xf6\x2c\x05\xd9\x97\x16\x21\x4c\xfa\x11\x72\xc2\x0d\xfc\x03\xbb\xc2\xca\xc2\x0d\x75\x8a\xdd\x22\x51\x10\xd3\x1a\x64\x1e\x2c\x96\xc2\xd0\x17\x86\xee\x2f\x78\x75\xf0\x36\x4e\xf9\xc3\x21\x4b\x88\xe5\x53\xf5\x5c\x69\x2f\x29\x8b\xee\x34\x64\xa7\xda\x8a\x80\xac\x75\xde\x07\xf0\x45\xeb\x9b\x2c\xc4\x8a\x2c\x15\xeb\x92\x2f\xa4\xdb\x40\x4c\xd7\xeb\x90\xb6\x1b\x8a\x60\x82\xbd\x8a\xaa\x1d\x2f\x25\x90\xcd\x57\xd1\x18\xf9\x44\x9f\x43\x31\xda\xbf\x90\xdd\x4e\xe2\x73\x8d\x70\xa7\xd6\xc3\xb3\xec\xa8\x48\x99\xf8\x4d\x01\x2f\xe0\x88\x5d\x2a\x10\x51\xa7\x23\x84\x81\xc3\xaf\xd5\x2c\x97\x29\x59\x99\x24\x24\x6e\xc4\xa1\xd1\xd6\xc6\x4a\x48\xd7\x71\x6d\x9d\x9f\x98\xf2\x11\xd3\xe1\x73\xe0\x39\xf2\xca\xc1\x6f\x51\xe5\x4e\x96\xb9\xcf\x1a\xfd\xe1\xa1\x4f\x21\x22\xf1\xc8\xd9\x7c\x45\xdf\xbb\x57\x06\x71\xed\x2e\xee\x05\x48\xe7\x4f\x54\xa9\xad\x95\x53\x7f\x0a\x58\x20\x91\x11\x8f\xb5\x11\xcf\x94\xe2\x92\x5a\xd3\x99\x88\x83\x43\x18\x38\x61\x34\x07\x49\xcf\x09\xc2\x34\x55\x8e\x67\x48\x92\x96\x85\xec\x22\xc7\x63\x32\x6c\xe8\x8f\xf6\x7e\x2f\x90\xf0\x8f\x3e\x6a\x11\xec\x6e\xc9\xc4\xbf\xf5\xf9\x16\x42\x66\x06\x8f\x49\x58\x58\xab\x53\xc9\xa0\x8f\x53\x7c\x15\x89\xdb\x17\x3e\x33\x7f\x96\xe4\x85\x69\xee\x54\x70\x33\xbb\xf3\x2e\x79\x68\x90\x41\x2e\x15\x82\x30\xf3\x8a\x93\x62\x12\xa1\x99\x6f\xb7\xed\x78\x91\xe1\x5c\x40\xe9\x49\x8c\xcf\x2c\x48\x1e\x3c\x72\x02\x45\x4b\xdc\x7c\x33\xaa\x96\xb8\xb9\x62\x58\x50\x0a\x69\x0e\xc8\xdb\x1d\x66\xfb\x8e\x5f\x44\x51\x52\xb0\x5b\x83\x5b\xe2\x66\x14\x0f\x21\xc0\x1a\xbe\xfa\xd3\xc5\xc0\xf7\x11\xe5\x0f\x6c\x83\x03\x3c\x7f\x2f\xc8\x3b\xae\xba\x14\x72\xe1\x0b\x92\xad\xf4\x32\x2a\x47\xfd\xc0\x05\xfc\x6a\x4e\x32\x1a\x10\x43\xb5\x07\xfc\x57\x25\x0d\xd7\xb6\xca\xca\xd9\x51\x5a\x72\x1f\xd6\xf8\x54\xc6\x9f\x96\x1d\xad\xb0\x80\x2b\x54\x20\x66\x0e\x0d\x88\xb2\xcc\xb9\x7f\xc2\x17\x1b\x4a\xed\xe1\x84\x5e\x2a\xaa\xd5\x04\x56\xc2\x48\x31\xcd\xb1\x51\x78\x8b\xae\x86\xb8\x3b\x25\x1e\x60\x9f\x45\x35\xf7\xa6\x8e\x3d\x6f\xf1\x4f\x89\x4c\x78\xf0\xc3\x9b\x3d\xd3\x79\xae\xd7\x9e\x1a\xa2\x9d\xe5\xe9\x3f\x6e\xb7\xc3\xd9\xd7\x5c\x38\x5c\x8b\x4d\x42\x49\x0f\x77\x8c\x87\x12\x8b\xb7\xb7\xf0\x8b\x5f\xc3\x89\x52\x53\xe0\x12\xa5\xa4\x1f\x62\x8d\xe9\x48\xb8\xce\x53\xeb\x2b\x62\xf1\xc6\xfe\x7e\x94\x14\x52\x0e\x83\x84\x74\x15\x10\xd4\x95\xe2\x3d\x18\x23\xeb\x50\xa5\x30\x16\x4d\xef\x4b\xd3\xa6\x42\x62\xd0\x19\x89\x6c\xf0\x42\x61\xa4\xd6\xd0\x80\x6d\x1f\xdd\x6f\xd7\xf7\x6f\x6e\xdf\xfc\x32\xbe\x14\x1f\x17\x9c\x56\x8c\x5f\x0b\xa3\xea\x7e\x3f\x51\xd9\x95\x95\xdf\xd3\x18\xa9\xc4\x87\xd8\xe8\xff\x14\x54\x97\x25\xf0\xcc\x57\x47\x88\xa3\x4f\x7d\x19\x5c\xc0\xc7\x17\x9f\x4e\xae\x87\xb4\xef\x49\xb7\xea\x9f\x90\xa1\x1b\xce\x1d\x19\x33\x19\xd1\x0c\x4b\x83\x29\xb9\x96\xc4\x60\x99\x8b\xb4\x33\xb9\x7a\x58\x78\x3c\x3a\xcf\x42\xa5\x97\xef\x99\xf9\xd8\x79\xf7\x82\x03\x3f\xfe\xb4\x5a\x2b\xca\xfa\x1b\x0c\xb5\x69\xad\xac\x8f\xcd\xb9\x45\x85\xeb\x1d\x70\xd6\xa1\x18\x49\x7b\x90\xc4\x39\x45\x6a\xbb\xd0\x55\x9e\x11\x79\x14\x2a\xc3\x7b\xeb\xbb\xb5\xbe\x95\xe4\xad\x3d\xcd\xe6\x4f\xc3\xd7\x34\x6a\x8a\x78\xfe\xc0\x56\x12\x5d\x1e\x03\x59\x97\xc3\xe2\x39\x9d\x55\x7f\xac\x4f\x40\xc9\xd9\xb1\x58\xf5\x6e\xde\x10\x52\x5e\x1f\x37\x34\xb6\x05\xe3\x6b\xb8\xf6\x33\xb8\x61\xc2\x72\x59\x48\x97\xc8\xb9\xd2\xa6\x93\xa4\xa8\xd2\x21\x5a\xe7\x25\x3e\xfb\xa3\x4f\xfb\x05\x72\xb2\x76\x1e\xdc\x58\xec\xe9\x42\xa8\x39\x92\x07\xe8\x20\xe0\x55\x8d\xb1\xae\xc8\xdb\xc8\x77\xbe\xf1\x1d\xe1\x1a\xc6\x04\x6e\x09\xbd\x54\xf3\x31\xba\xc0\x14\xd8\x24\xd7\xf3\xc4\xca\x3f\xba\x08\xc8\xf5\xfc\x9d\xfc\x83\x0b\x3f\x7e\xc1\x0e\xc7\x8d\x8a\x0a\xc5\x8e\x90\x82\xe8\xf8\x2c\xf0\x47\xce\x5e\x7e\xfa\x71\x34\x29\x05\x16\xda\x6c\xfa\xa8\xf1\x33\xce\x25\xe8\xa7\xbf\xfe\x17\x93\xf4\x9f\x3f\xfd\x75\x34\x4d\x64\xfb\x75\xd5\x55\x51\x0f\xa3\x67\x11\xf3\xa3\x97\xcf\x7f\xfc\x48\xff\x86\xe9\xe1\xe6\x68\x52\x1a\x5d\xa2\x71\xb2\x33\xa9\x88\x16\xb0\x65\xaf\x7c\x4b\xdd\xbb\xb3\xd0\x54\xf7\x9d\xd6\x06\x58\x6c\xbe\x1f\xb7\x89\xd1\x24\x66\x9a\x15\x8e\x2c\xa3\x74\xa0\x2b\x67\x65\xc6\x1b\xf1\x60\xc4\x4a\x5a\x98\x56\x32\xcf\xfa\x3b\xb3\xcc\x8a\x37\x07\x86\xd4\x76\x94\x29\xa8\xb5\x7f\xc7\x20\xa8\x3d\x83\x1e\xa4\xcd\xfd\x66\xca\x7e\xfc\xaf\x51\xdc\x8f\x8f\x93\x42\xaa\xd0\x7d\xa4\x2f\x22\x1d\xe8\x65\x30\xa9\x31\x76\xf0\x87\xac\xcb\x4c\xc4\xfe\x50\x98\x45\xf1\xc3\x5e\xab\xe8\x48\x39\xb9\xb3\x1b\x74\x56\x0b\x88\xa9\x0d\x0d\x66\x2e\x59\xf4\xd6\xdc\x0e\x7a\x87\x3b\x26\x66\xaf\x18\xd7\x44\xaf\x39\xbf\x9c\x53\xda\x2d\x42\xae\x3c\x4c\x52\xcc\x81\x07\xdb\xa7\x0f\x07\xd5\xad\x76\xc0\x10\x9e\x17\x60\x06\x4a\x8f\xbb\x03\xc0\xd8\x5b\xd7\x6f\x58\x28\x63\x88\x38\x7a\x39\x25\x78\x9c\xfd\x28\x7c\x1d\x7a\x54\xbe\xd3\x7b\xac\x46\x37\x42\x42\xad\x47\x42\x89\x5e\xa1\x31\x32\xcb\xb0\xab\xd2\x44\x14\xb6\xdf\x0c\x35\xd7\xa7\x9a\xa5\x31\x56\x68\xdf\x8e\x19\xbb\x51\x89\xb4\x49\x59\x4d\x73\xd9\xf5\x2e\xdb\xef\x0a\xcf\x8d\x9d\x16\xff\x2c\x8a\x62\x7b\x5e\x78\x90\xc5\x5f\x90\xb9\x60\xdb\x32\x45\x58\x49\x5f\x50\xa0\x73\x98\x0a\xb6\x34\xfe\x5e\x3c\x66\x30\xdd\x80\x50\x1b\xad\x7a\x9e\x19\x31\xad\xb1\x30\x88\xd3\xf0\xf8\x73\xc0\x8d\x1f\xd6\x05\xb9\xe5\xc1\x8d\x17\x95\xd1\xff\x97\xe1\x9d\xe6\x7e\xcf\x83\x0e\x02\xff\xa1\x0d\x9c\x5e\x78\xe7\x1e\xbe\x85\x05\x3d\xa9\x87\xa7\xb4\xd5\xdb\x22\x72\x7b\xab\x88\x5d\x1d\x0f\xd2\xb0\x76\x9b\x68\x54\x03\x2b\xfc\x81\x90\x7a\xd1\x04\x5e\x68\xb5\x22\x73\x1f\x52\x82\x06\x85\xd3\x3b\xe0\x87\x55\x76\x9f\xab\x81\x5e\x5d\x5f\x75\xb4\xe1\x2d\x0e\x9c\xc8\x5d\xdd\x32\xdb\xe7\xaf\x8d\xa8\xe6\x70\x54\x83\xad\xe6\x31\x56\x42\x0c\xda\x52\x2b\x8b\x7d\x57\x9e\xf6\x88\xe6\x1a\xd8\x7e\xae\x1b\xc6\x63\x56\xdb\xca\x92\x63\xbd\xa2\xae\xb3\x2d\x9c\x2b\xfd\x1f\xe3\xf1\xa8\xd9\xaf\x4d\xe0\x05\x79\x18\xbe\x23\xd1\xfe\xdd\x3b\x75\x76\x39\xe1\xe7\xc0\x34\x43\x21\x7f\xd2\x50\xd6\xa1\xb1\x2f\x6f\xfe\xf7\xfd\x2f\xa3\x53\x57\x9e\x7d\x5a\xde\x9a\x4d\xe7\x89\x45\x61\xd2\x05\x69\x4d\x34\x7a\x75\x63\xab\x53\x75\xc2\x8a\xda\xe8\xed\xb6\xc2\xa2\x08\x23\x8f\x3e\x38\x18\x08\x7f\x89\x94\x7d\xcf\xf0\xad\xbd\xc2\x99\x1e\x81\x48\xab\x5d\xa6\xbf\x5a\xd9\xf3\xf7\x49\x5e\x1e\xb9\xdf\x13\x24\xf2\x0c\x7e\x66\x0a\x9a\x3f\x87\xc1\x65\x5e\x02\x76\x2a\x01\xfd\x0f\x3a\x4f\xa7\xa1\x7d\x7b\x33\xde\x36\x3e\xed\x91\xde\xde\xeb\xab\x93\x1f\x5e\x1d\x7f\x4e\xf7\xe4\xd3\x93\xff\x0b\x00\x00\xff\xff\xb9\x73\x67\x54\xa2\x4c\x00\x00")

func wski18nResourcesEn_usAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_wski18nResourcesEn_usAllJson,
		"wski18n/resources/en_US.all.json",
	)
}

func wski18nResourcesEn_usAllJson() (*asset, error) {
	bytes, err := wski18nResourcesEn_usAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wski18n/resources/en_US.all.json", size: 19618, mode: os.FileMode(436), modTime: time.Unix(1536397293, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wski18nResourcesEs_esAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesEs_esAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_wski18nResourcesEs_esAllJson,
		"wski18n/resources/es_ES.all.json",
	)
}

func wski18nResourcesEs_esAllJson() (*asset, error) {
	bytes, err := wski18nResourcesEs_esAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wski18n/resources/es_ES.all.json", size: 0, mode: os.FileMode(436), modTime: time.Unix(1536397114, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wski18nResourcesFr_frAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8a\xe6\x52\x50\xa8\xe6\x52\x50\x50\x50\x50\xca\x4c\x51\xb2\x52\x50\x4a\xaa\x2c\x48\x2c\x2e\x56\x48\x4e\x2d\x2a\xc9\x4c\xcb\x4c\x4e\x2c\x49\x55\x48\xce\x48\x4d\xce\xce\xcc\x4b\x57\xd2\x81\x28\x2c\x29\x4a\xcc\x2b\xce\x49\x2c\xc9\xcc\xcf\x03\xe9\x08\xce\xcf\x4d\x55\x40\x12\x53\xc8\xcc\x53\x70\x2b\x4a\xcd\x4b\xce\x50\xe2\x52\x50\xa8\xe5\x8a\xe5\x02\x04\x00\x00\xff\xff\x45\xa4\xe9\x62\x65\x00\x00\x00")

func wski18nResourcesFr_frAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_wski18nResourcesFr_frAllJson,
		"wski18n/resources/fr_FR.all.json",
	)
}

func wski18nResourcesFr_frAllJson() (*asset, error) {
	bytes, err := wski18nResourcesFr_frAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wski18n/resources/fr_FR.all.json", size: 101, mode: os.FileMode(436), modTime: time.Unix(1536397114, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wski18nResourcesIt_itAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesIt_itAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_wski18nResourcesIt_itAllJson,
		"wski18n/resources/it_IT.all.json",
	)
}

func wski18nResourcesIt_itAllJson() (*asset, error) {
	bytes, err := wski18nResourcesIt_itAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wski18n/resources/it_IT.all.json", size: 0, mode: os.FileMode(436), modTime: time.Unix(1536397114, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wski18nResourcesJa_jaAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesJa_jaAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_wski18nResourcesJa_jaAllJson,
		"wski18n/resources/ja_JA.all.json",
	)
}

func wski18nResourcesJa_jaAllJson() (*asset, error) {
	bytes, err := wski18nResourcesJa_jaAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wski18n/resources/ja_JA.all.json", size: 0, mode: os.FileMode(436), modTime: time.Unix(1536397114, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wski18nResourcesKo_krAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesKo_krAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_wski18nResourcesKo_krAllJson,
		"wski18n/resources/ko_KR.all.json",
	)
}

func wski18nResourcesKo_krAllJson() (*asset, error) {
	bytes, err := wski18nResourcesKo_krAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wski18n/resources/ko_KR.all.json", size: 0, mode: os.FileMode(436), modTime: time.Unix(1536397114, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wski18nResourcesPt_brAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesPt_brAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_wski18nResourcesPt_brAllJson,
		"wski18n/resources/pt_BR.all.json",
	)
}

func wski18nResourcesPt_brAllJson() (*asset, error) {
	bytes, err := wski18nResourcesPt_brAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wski18n/resources/pt_BR.all.json", size: 0, mode: os.FileMode(436), modTime: time.Unix(1536397114, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wski18nResourcesZh_hansAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesZh_hansAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_wski18nResourcesZh_hansAllJson,
		"wski18n/resources/zh_Hans.all.json",
	)
}

func wski18nResourcesZh_hansAllJson() (*asset, error) {
	bytes, err := wski18nResourcesZh_hansAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wski18n/resources/zh_Hans.all.json", size: 0, mode: os.FileMode(436), modTime: time.Unix(1536397114, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wski18nResourcesZh_hantAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesZh_hantAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_wski18nResourcesZh_hantAllJson,
		"wski18n/resources/zh_Hant.all.json",
	)
}

func wski18nResourcesZh_hantAllJson() (*asset, error) {
	bytes, err := wski18nResourcesZh_hantAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wski18n/resources/zh_Hant.all.json", size: 0, mode: os.FileMode(436), modTime: time.Unix(1536397114, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"wski18n/resources/de_DE.all.json":   wski18nResourcesDe_deAllJson,
	"wski18n/resources/en_US.all.json":   wski18nResourcesEn_usAllJson,
	"wski18n/resources/es_ES.all.json":   wski18nResourcesEs_esAllJson,
	"wski18n/resources/fr_FR.all.json":   wski18nResourcesFr_frAllJson,
	"wski18n/resources/it_IT.all.json":   wski18nResourcesIt_itAllJson,
	"wski18n/resources/ja_JA.all.json":   wski18nResourcesJa_jaAllJson,
	"wski18n/resources/ko_KR.all.json":   wski18nResourcesKo_krAllJson,
	"wski18n/resources/pt_BR.all.json":   wski18nResourcesPt_brAllJson,
	"wski18n/resources/zh_Hans.all.json": wski18nResourcesZh_hansAllJson,
	"wski18n/resources/zh_Hant.all.json": wski18nResourcesZh_hantAllJson,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"wski18n": &bintree{nil, map[string]*bintree{
		"resources": &bintree{nil, map[string]*bintree{
			"de_DE.all.json":   &bintree{wski18nResourcesDe_deAllJson, map[string]*bintree{}},
			"en_US.all.json":   &bintree{wski18nResourcesEn_usAllJson, map[string]*bintree{}},
			"es_ES.all.json":   &bintree{wski18nResourcesEs_esAllJson, map[string]*bintree{}},
			"fr_FR.all.json":   &bintree{wski18nResourcesFr_frAllJson, map[string]*bintree{}},
			"it_IT.all.json":   &bintree{wski18nResourcesIt_itAllJson, map[string]*bintree{}},
			"ja_JA.all.json":   &bintree{wski18nResourcesJa_jaAllJson, map[string]*bintree{}},
			"ko_KR.all.json":   &bintree{wski18nResourcesKo_krAllJson, map[string]*bintree{}},
			"pt_BR.all.json":   &bintree{wski18nResourcesPt_brAllJson, map[string]*bintree{}},
			"zh_Hans.all.json": &bintree{wski18nResourcesZh_hansAllJson, map[string]*bintree{}},
			"zh_Hant.all.json": &bintree{wski18nResourcesZh_hantAllJson, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
