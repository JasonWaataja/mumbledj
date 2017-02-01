// Code generated by go-bindata.
// sources:
// config.yaml
// DO NOT EDIT!

package main

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

var _configYaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5b\xdd\x8f\x1b\x37\x92\x7f\xd7\x5f\x51\x91\xcf\x38\x1b\x98\x95\x3f\xb2\xd9\x5d\x08\x5e\x07\x13\xdb\x7b\xf1\xc1\x4e\x02\xdb\x59\x60\x9f\x1a\x54\x77\xb5\x9a\x3b\xdd\x64\x2f\x3f\xa4\x51\x1e\xee\x6f\x3f\x54\xf1\xa3\x5b\x6a\xcd\x48\x1a\x07\x38\x1c\xb0\x7e\x49\xc4\x2e\xfe\x8a\xac\x2f\x56\x15\x39\x8f\xe0\xa3\xef\x56\x2d\xbe\xfd\xef\xd9\x23\xf8\x61\x07\x1f\x85\x73\x8d\x44\x0f\xff\x65\x24\xae\xd1\xcc\x1e\xc1\x1b\xdd\xef\x8c\x5c\x37\x0e\x9e\x94\x4f\xe1\xe5\xf3\x17\x7f\x9a\x50\xc1\x93\x8f\xef\xbf\xc0\x07\x59\xa2\xb2\xf8\x74\xf6\x08\x4a\xad\x6a\xb9\x5e\xec\x44\xd7\xce\x66\xa2\x97\xc5\x0d\xee\xec\x72\x36\x03\x00\x78\x04\xff\xd0\xfe\x8b\x5f\x21\x5c\xff\xf2\x1e\x6e\x70\xb7\xe0\xe1\x9d\xf6\xce\xaf\x70\x09\xf3\x79\xa2\xfb\xac\xbd\xaa\xde\xb4\xda\x57\xfb\xa4\x8f\xe0\xa7\x9f\xbf\xbc\x5b\xc2\x97\x26\x63\x80\xb4\x84\x60\xa0\x6c\x25\x2a\x07\xef\xdf\x06\x52\x4b\x10\x25\x41\x04\xe0\x59\x85\xb5\xf0\xad\x1b\x16\xf3\x36\x0c\x40\xa9\xbb\x8e\x66\x3a\x0d\x2b\x04\xd1\xf7\xad\xc4\x8a\x7f\x69\xb7\xcf\xf6\x7d\x4d\xac\xa0\xd2\xa0\xb4\x83\xad\x50\x0e\x44\x9e\xbe\xda\x41\x64\x71\x05\x16\x19\x0e\xbb\xde\xed\xc0\x3a\x23\xd5\x1a\x9e\xcc\xe7\x4f\x03\x5c\x9c\xb1\x84\xf9\x8f\xd8\xb6\xfa\x1b\x78\x0f\xa2\x03\xc1\xfc\xe0\xcb\xae\x47\xf8\xa6\xc1\xb6\x87\x5a\x1b\x10\xd0\x4a\xeb\x40\xd7\x3c\x4b\xa8\xca\x2e\xe6\x93\x0d\x34\x42\x29\x6c\x99\xde\x35\x48\x38\xcc\x5d\x39\x34\xe0\x7b\xad\x48\x2b\x0a\x4b\x27\xb5\x3a\xba\xa1\xad\xb4\xcd\xe1\xec\x38\x85\xfe\x97\x46\x8d\xd6\x99\xd1\xc9\xfd\x05\xb2\xb1\x42\xdf\x84\xc5\xd3\x24\x6f\x91\xfe\xd3\xb7\x62\x07\xc2\x57\x52\x43\x2d\x5b\xb4\x0b\x56\xaa\xdb\x6a\xb0\xbe\xef\xb5\x71\x58\x41\xd9\x68\x59\xa2\x05\x61\x10\xe6\x75\xdd\xf5\xb8\x9e\x03\xc1\xcc\xc5\xa6\xd4\x6a\x33\x0f\xfc\x08\x0a\x4d\x11\x05\xb4\xcc\xa4\xb3\xd9\xec\x5f\x1e\x3d\x66\x8d\x7f\x12\x4e\xd2\x76\x84\x83\xce\x5b\x47\xea\xee\xd0\x81\x36\x80\xb7\x25\x62\x15\xd4\xee\x8c\x5c\x93\x69\x0b\x70\x46\x94\x37\x60\x6f\x64\x1f\x18\xf1\xef\x82\x7e\x17\x86\xa0\x96\xf0\x7c\xf1\xdd\x43\xc1\x69\xd5\xac\xdb\x01\x3f\x0d\xdd\xc5\xe2\xa3\xb8\x95\x9d\xef\xe2\xba\x2a\xcf\x14\x0a\xa4\x02\x8b\xa5\x26\xdb\x80\xcf\x41\x33\xcf\x59\x9d\x5e\x19\x24\xed\x94\x24\xcc\x44\x1e\x58\x75\xe2\xb6\x08\xdb\x49\xe3\x4b\x78\x7e\x94\x8f\x85\x1e\x4d\x5e\xda\x7d\x1c\x12\x8d\x3d\x60\x61\x8b\x1e\x4d\x91\xbe\x2e\xe1\xbb\xcc\xe8\xbd\x05\xdb\xf8\xba\x6e\xc9\x80\x50\x89\x55\x8b\x15\x6c\x1b\x54\xd9\x12\xad\x13\xc6\xd9\xef\x99\x5e\x78\xa7\x3b\xe1\x64\x59\x84\x49\x58\xd0\xaa\x6b\xd1\x5a\x4c\x80\xd7\x4a\x69\xaf\x4a\x8c\x22\x92\xaa\xd6\xa6\x0b\x52\x12\x2e\x80\xe2\x5a\x2a\x45\xfc\x74\x1d\xed\x8f\x56\xb6\x12\xe5\x4d\xe4\x12\x21\x0a\x85\xdb\xb8\xfe\x25\x38\xe3\x71\x36\x9b\x0d\x7e\x94\x6d\xea\xba\xaa\x0c\x5a\x1b\x16\xdb\x68\xdf\x56\x20\x9c\x23\xcf\xd8\xf7\xa2\x20\x13\x11\xa8\x97\x30\x7f\xf1\xf2\xcf\x8b\xe7\x8b\xe7\x8b\x17\xd9\x47\x7e\xd1\xc6\x9d\x09\x43\xfe\xb1\x84\xf9\x9f\xfe\xf8\xe7\x6f\xff\x32\xcc\x17\xd6\x6e\xb5\xa9\x58\x31\x69\xa5\x6a\x4d\xf3\x2d\x9a\x0d\x9a\x89\xef\x2b\x0d\x7d\x9c\x74\xca\xa7\x13\xdd\xd8\xa9\x7f\xb5\x68\x94\xe8\x90\x19\xa6\xd3\x24\x90\xfb\xf8\x69\x09\xf3\xf4\x21\x4f\xfb\x9b\x6c\xb1\x17\xae\x89\xc1\xc0\x40\xff\xe2\x25\xc7\x80\x10\xf0\xbc\x6b\x50\x39\x59\x0a\x5e\xbc\xb0\x20\xc0\xe0\x5a\x5a\x87\x06\x2b\x9e\x70\x74\x1f\x09\x43\x5a\x50\xec\x6d\xa7\x76\x44\x48\x45\xff\xe2\xe5\xde\xb9\x13\x24\x9f\x8c\x2f\x69\x40\x90\x8f\x59\x2c\xbd\xc1\x51\x28\xfd\x3e\xdb\xdc\xb1\xaf\x50\x69\xb4\x7c\x4e\x6c\xd0\xc8\x7a\xc7\xa0\x25\x1a\x27\x6b\xda\x1b\x92\xf9\xd1\x50\x50\x0d\x6d\x3d\xc2\x95\x5a\x59\xda\xad\x2a\x77\x0b\x78\xef\x68\x43\x2b\xb4\xbc\x93\x16\xc5\x06\xc1\x35\xd2\x82\x56\x57\xb0\xf2\x0e\x2a\x69\xc9\x6b\x40\x3a\x90\x21\x98\x53\xb0\x6c\xc4\x46\xaa\x75\x04\x94\xd6\x7a\xb4\x07\x16\x21\x12\x63\x12\xb9\x41\x30\x3e\x38\x45\xe7\x5b\x27\x7b\x02\x54\xd6\x09\x45\xd1\x57\xd7\x07\xca\x4d\xbb\x3d\xf0\xbd\xb1\x5e\xc7\x1b\x25\xb5\x1c\x53\xd9\x21\xcd\xf9\xaa\xa3\x99\x63\xb5\xdd\xc5\x99\xd2\x83\xbb\xb8\xc7\xd4\xe1\x3c\x86\x37\xb8\x1b\xf3\xbb\x2e\x4b\x72\x79\xa7\x6f\x50\xd1\x7f\x40\x2a\xe9\xa4\x68\xe5\x6f\x98\x6d\x67\x2b\x5d\x43\xb0\xbd\x30\x82\x82\xe3\x6a\x17\x4e\x70\x7b\x6c\x31\x62\x0f\x90\xf4\x71\xde\xba\xc2\xbc\x22\xcc\xbb\xcf\x90\x53\xe4\x14\x6d\xbb\x1b\x07\x16\x83\xce\xec\xc6\x56\x3b\x36\x0d\x51\x53\x02\x51\x49\x3b\x98\x4e\xb0\x79\x9e\x55\xc4\x78\x9d\x82\x63\x60\xfc\xa3\xde\x42\x27\xd4\x0e\x9c\xec\xd0\xa6\x50\x76\xe8\x50\xcc\xf9\x20\xc3\x08\x4c\xc7\x0c\x22\xb5\x5d\xc2\x8b\xe7\x13\xfc\x78\xe2\x1d\x72\xd8\x0a\xf2\x04\xf5\x87\x15\xba\x2d\xe2\x38\xf3\x89\x7b\x4d\xa0\x63\x46\x92\x32\xa5\x8d\x68\x97\xf0\x1d\x05\x79\x51\x36\x43\xce\xf0\x86\x7e\x81\xd5\x6a\x6d\x29\x18\xb9\x06\x77\xac\xa0\x4a\x6f\x55\xab\x45\x85\x55\x40\xca\xd2\xd8\xf3\x89\x7c\x92\x6a\x27\xda\x60\xe5\x96\xac\x84\xf2\x39\x06\xae\xa4\xc1\xd2\x69\xb3\xa3\x53\xfc\xa3\xfc\x21\x1f\x9d\x34\xad\x20\xda\x25\x7c\xf7\xe2\x65\x8e\xf1\x68\xa4\xae\x38\x76\xc8\x0e\x43\xb6\x11\x25\x80\xad\xe8\x2d\x9d\x6f\xb5\x36\x48\x4a\xd4\x6a\x4d\x16\x5e\xb6\x28\x28\x72\xd6\x46\x77\x21\x08\x11\xe3\x2b\xe2\xd7\x68\x6f\xa2\x3d\xe2\x6d\x2f\x0d\x16\x84\xba\x84\x97\x7f\xbc\x83\x5f\x92\x2a\x8a\xb2\x81\xb2\xc1\xf2\x26\x85\xb1\xb0\x9b\x9a\xb3\x1d\x42\xaa\x40\x3a\xec\x2c\xb3\xe9\xa4\xf2\x0e\x6d\x4a\x0e\xb1\xbc\xd9\x97\x78\xcc\x66\xb3\x24\xe8\xc0\x72\xb4\x09\x06\x8d\x48\x0b\x78\xa7\x36\xd2\x68\xc5\xc9\xf6\x46\x18\x49\xf2\x0e\xce\xc2\x11\x30\xa4\xef\xde\x62\x05\x0d\x9a\xe8\xf3\x59\xbc\x4b\x98\xff\xc7\x8f\x3f\x7f\x7c\xf7\x6c\xc1\xa0\xcf\x3a\x8e\x68\xd5\x3f\x29\x49\xdc\xe8\xd6\x77\x38\xa9\x0b\xc2\x70\xc4\x09\x63\x94\x8d\x65\x5d\x7c\xd0\x5b\x8a\xcb\x81\x0c\x44\xdb\xea\x2d\x56\x81\xbc\xe5\x4f\x44\xfd\xfc\x45\xb6\x5c\xb9\x6e\xee\xa2\x6f\xc2\x37\x9a\xf0\x97\xd9\x6c\x26\xaa\x4e\xaa\xa1\x50\x79\xc7\xa6\x05\x61\xf4\xfb\xc3\xf0\xc1\xc7\x81\xb4\x29\x52\xb0\xf9\x5d\x01\xb9\x48\x4c\x88\xa1\x14\x8a\x44\x83\xb7\x58\xfa\x18\x8a\xe8\xf3\x70\x94\x1e\xf5\xe4\x0f\xb1\xee\x60\xb6\x40\x87\xf9\x61\xe8\xe2\xb3\x89\xfc\x98\xca\x19\xce\x6f\x9b\x98\x64\x31\x35\xa9\x9e\x17\xc7\x59\x2e\x1f\x32\xc3\x39\xae\x39\xc5\x8b\x78\x31\xde\xd8\x98\x3e\xcb\xae\xd7\x44\x66\x69\xe5\x74\x82\xc6\x95\xc7\xa5\xe4\x42\x88\x67\x33\xab\x25\xff\x2f\xfd\xfb\x03\xcc\x3f\xfb\x1e\x0d\xe5\x26\xa4\x5b\xae\x2d\x96\xc7\x6d\x0c\x85\x29\x43\xdd\xd3\x79\x2b\x4b\x90\x6a\x01\xd7\x41\x2d\xf4\xbd\xa1\xc3\xd6\x36\xd8\xb6\x64\xd3\x42\x55\x24\xa4\x7c\x48\x77\xbd\x56\xa8\x72\xae\x4b\x00\xc5\xd8\xda\xfe\xe7\xd9\x47\x1a\x9b\x87\xcf\x68\xad\x58\x8f\xd7\xa9\x74\x41\x4b\x2b\x6a\x2a\x53\x0b\x34\x46\x9b\x25\xcc\x7f\xd2\xd0\x09\x57\x36\x14\xe5\x63\x46\xe4\x55\xc5\x3e\xc4\x2b\xcc\x0c\x16\xf3\x11\x92\x2a\x32\xfb\xa2\x37\x58\xcb\xdb\x0c\x78\x1d\xa2\x5d\x3c\xf4\xa9\xf4\x62\x58\xed\x9d\x95\x15\x87\xa1\x29\x30\xa5\xb9\x41\xc2\x43\x10\x6c\x84\x11\x25\x97\x93\x36\x08\xa7\x42\x2b\xd7\x8a\x0e\xd9\x44\x1c\x02\x8c\xa2\x7c\xbb\x05\x87\xb7\x2e\xef\x7a\xdf\x6c\x7e\x56\xed\x0e\xb4\x42\x2a\x13\x23\xe8\x13\xb2\x99\x5a\x1a\xeb\x9e\x92\x49\x11\x8f\x98\x75\xf2\x6e\x96\x30\xff\x26\x9e\x6c\xc4\x8c\xf6\x7b\x4c\x9e\xb1\xd6\x48\x5b\xff\x42\x51\x20\x9c\xa5\x3a\x55\x32\x32\x14\x16\x5c\x15\x8e\x44\x48\x21\x41\xaa\x75\x11\x93\xa9\x41\x1f\x6f\xc2\x07\x0e\xa3\xde\x18\x54\xae\xdd\xa5\x94\xab\x1a\x0a\xf1\x1f\xb0\xd5\x5b\x22\x1a\xaa\x75\x8e\x89\x49\x32\x43\x45\xbb\xda\x0d\xb9\x14\xbc\xe3\x28\x1a\x9d\xb4\x11\x36\xa2\xb9\xc6\x20\xc6\x46\x8a\x37\xec\xfa\xba\xa7\x13\x2c\x6e\xf7\x11\x88\x56\x0a\x8b\x76\x09\xd7\x99\x5f\xf0\x38\x76\x9f\xe8\xee\x49\x53\xc9\x79\x46\x2b\x5a\xe4\xcc\xb0\x60\x97\x0a\x8e\x0f\x7f\x05\x4d\xba\x09\x71\x86\x61\x8e\xcc\xbd\x0a\x11\x06\xfe\x4a\x31\x84\xd5\x78\x9c\x2e\xf1\xa8\xd0\x96\x46\xf6\xa1\xc2\x7c\x3b\xfc\xa0\x63\x6b\xab\x72\xd7\x21\x89\x21\x17\x7f\xdc\x01\x49\xa3\xd2\xe6\xe8\x95\x70\xb3\x09\xc0\xdf\x85\x91\xda\xdb\x3c\x12\x6b\x70\xb1\x23\x29\x58\x3a\x2d\xb8\x88\x18\x9b\xe4\xe8\x30\x8c\xab\xa5\x62\xa6\xf6\xb1\x87\x62\x84\xb2\x2d\xd7\x1f\x91\xd9\xf0\x2f\xa4\x60\x9c\xf4\x69\xd7\xa0\x81\x56\xa8\xb5\x67\x2b\x87\xb7\x9a\x4c\x1c\x0c\x76\x9a\x22\x47\xa2\xa4\xd5\x70\xd5\xc9\x39\x21\xcc\x1f\xcf\xe1\x89\xf5\x65\x43\xcb\x9a\x3f\xb6\xf3\x2b\x98\x3f\xae\xe6\x57\x80\xae\x5c\x3c\x9d\x30\x4c\x39\x87\xf5\x2b\xeb\xa4\xe3\x00\xce\x38\xc6\x2b\x3e\x93\x2b\xe1\xc4\x02\x3e\x11\x53\x76\xf3\x06\xed\xc0\x7c\x2b\xdb\x16\x4a\xc1\x3d\x97\xa1\xb7\xd3\x49\xbb\x42\x0a\x6f\xb9\x28\x1d\x1c\x29\xd9\xd6\x6c\xb4\x06\x8a\xaa\xa2\xaa\xe6\x93\xb1\x61\x64\x30\xa5\x90\xff\xa4\xf1\x3d\xf5\xcf\xaf\xab\xca\xe6\xbe\x8a\x1e\xba\x0a\x41\x1f\x02\x3a\xac\xa4\x00\x2b\x1d\xa6\xac\xf0\xd0\x55\xa7\x9e\x1f\xbd\xdf\x9b\x76\x88\x7a\xf0\xeb\xa7\x0f\xb9\x0b\x43\xde\xc7\x2d\x3d\x16\x1b\x81\x8a\xaa\xca\x8a\x9f\x1f\x02\x6d\x44\x2b\xab\xc3\x60\xf2\x93\x06\x1e\x4f\x81\x64\x4b\xb1\x25\x04\xe7\x8c\xda\x1b\xbd\x91\x74\x0c\xfe\xfa\xe9\xc3\x13\xfb\xf4\x00\x39\x02\x3a\xad\x8b\x56\xab\x75\x46\xfe\x87\xf6\x26\x7c\x7c\x62\x9f\x06\x5c\x94\x6c\x59\x4e\x6b\x20\x52\x2e\xd1\x14\xf0\x04\xd0\x25\x07\x22\x72\x14\x0a\xe4\xbd\xd1\x54\x01\x44\xc5\x77\x0b\xf8\x49\x0f\x60\x7c\x80\xad\x28\x5f\x13\x55\x85\x87\x5b\xd5\x0a\x63\x07\x88\xbf\x2e\x61\xfe\x6a\xf5\xfa\xb1\x7d\xf5\x6c\xf5\x3a\xd0\xc3\xab\xd5\xeb\x17\xfc\x33\xe8\x6b\xac\x91\xe5\xab\x95\x79\xfd\x4a\x32\xbd\x7c\x1d\xd4\xf7\xd8\xee\x33\xa0\x34\x3d\xc9\xf1\x1e\x16\x8f\xab\x81\x87\xbd\x4b\xed\xac\x1b\xdf\x15\x07\x52\x64\x44\xf3\x7a\x82\x52\x72\x46\x4c\xa9\xc3\x0a\x23\xa7\xca\xb3\x4d\x45\x29\x1a\x58\x61\x76\x8b\x90\xd0\x27\x71\xa7\xb0\x2e\xaa\xaa\xd5\xa5\x68\xcf\x72\x0d\xa6\x9c\xfa\x47\x7b\xcc\x41\x38\xc9\x7a\x88\x7f\x90\x60\x98\x51\xe8\x98\xee\xac\xc3\x6e\x4f\x60\x67\xb8\x89\x30\x6b\x4f\x19\xf4\xd8\xb2\xd3\xd8\x60\xc1\x54\x05\x1e\xdd\x94\xd2\x45\x4a\x50\x0a\x3b\x36\xe3\x71\xe6\x42\x1f\x26\xd3\x86\x25\x17\x94\xee\xc9\x12\xc7\x73\xff\x36\x6c\x28\x7e\x05\xb1\x11\xb2\xa5\xd3\x6f\x1f\xca\x2b\xca\x38\xd7\x4a\xfe\x86\x55\xe1\x76\xfd\x00\x43\x18\x34\x30\x74\x5e\x78\x45\xd0\xf5\xdf\x92\x2c\xbb\x6f\xfd\x57\xfb\x7a\x3a\xad\xb2\xa0\x78\x57\xff\x76\xf4\xff\xef\x8e\xae\xf0\xd6\x9d\xe5\xe7\x44\x38\x75\x73\xf5\xbb\xba\xf9\xde\x31\x18\x8e\x7f\x20\xbe\x5c\x07\xdf\x95\xc2\x3e\x4a\xdb\xa0\xbc\x2f\xcc\xc9\xf9\x50\x85\xb5\x54\x18\xac\x57\x54\xd5\x22\xa6\xd2\x54\x07\x73\x83\xe1\xe4\xc6\x33\xe9\x64\xeb\xa5\xbd\x70\xeb\x3f\x7b\xd7\x7b\x17\x16\xb8\xd7\x0e\x19\x9a\x08\xa1\x11\x02\xb2\x4e\x79\x39\xa7\x7f\x2a\xe5\xdc\xd9\xea\x8e\x86\xb8\x98\xa6\xc7\xce\x09\x15\x01\x69\xe8\x18\x27\xcb\x76\xb9\x78\xb9\x21\x8e\x64\x57\xc9\x26\xe2\x1c\xd6\xd0\x19\xf2\x19\x51\x4f\x45\x14\x3e\x4e\x83\xe2\xf0\xed\xd2\x34\x2a\x09\x71\xef\x7a\x63\xa5\x7d\x28\xbe\xd3\x7e\xd3\x15\xc8\x60\x2f\x24\x53\x4a\xd9\xf1\x96\x6f\x69\xce\x95\x65\x90\xc2\xbe\x30\x23\xb8\x85\x1c\x1b\xae\xa2\xff\xad\x76\x90\x9d\x3f\x89\xb3\xd6\xa6\x44\x7b\x23\xfb\xd3\xb2\xcc\xa4\x13\x61\xd5\x97\xda\xda\xfb\x8e\x1d\xc9\x61\xbb\xe3\x4b\x36\x3b\x15\xcf\x49\x19\x0c\x57\x7e\x3d\xc7\xb5\xa9\x0c\x1a\x61\x43\xec\xa5\x95\xcb\x55\xe4\xd5\x9f\x92\x44\xbe\x0e\x3b\x5f\x22\x69\xca\x11\xc9\xf4\xbf\xab\x68\xf2\x65\xdf\x19\x09\x45\xbe\xb3\x1c\xd5\xdd\x53\x2b\xa1\x08\xdd\x0b\x13\x1a\x4a\xc7\xf0\xe1\xf0\xfa\x73\x2a\xee\x1c\x25\x2f\x93\x38\x15\x92\xa7\x85\x4c\x54\x13\xb9\x36\x0f\x75\xcc\xa1\xed\xb5\x7f\x71\x7f\xbf\x34\x13\x61\xd1\xa0\xa8\xd0\x0c\x67\xde\x9b\xd4\x8f\xa1\x7d\xd1\xd8\xfe\x4a\x79\x61\xc5\x9d\xb3\xaf\xb9\x6b\x76\x04\x83\x41\xfe\xa9\xa5\xea\xce\x38\x03\x02\xdd\x44\x44\x34\x7c\xa1\xed\x7d\xd4\x1b\xb4\xb9\x0d\x02\x52\x39\x1d\x5f\x70\x44\x45\xa7\xf7\x0c\xb2\x0e\x76\xd3\x8a\x1d\x5f\x05\xf2\x65\xad\xd3\x60\x75\x87\x1c\xc6\x5a\x7b\xba\x34\xe4\x2a\xdd\x16\xc2\x60\xd1\xf2\xb5\x9a\x1c\xe5\x64\xbf\x5a\x34\xdc\x15\x12\x2a\x54\xf3\x89\x35\xe5\x09\x99\x9c\x4b\xe6\xc3\x24\x45\xaa\x82\x16\x5d\x0c\x8f\x1d\xf8\x15\x87\xd2\x5b\xc2\x0b\xfb\x09\x9f\x52\xef\xea\x46\xb6\x67\x14\x13\x44\x35\x91\xf2\xcd\x85\x22\xfe\xec\x74\x74\x69\xbe\x75\x51\x15\xdf\x02\x28\x0b\xd2\xd9\xc3\x8b\x87\xe4\x27\xb4\xdd\x78\xdd\x7d\x72\x91\x03\xed\x64\xa9\xfc\x00\x40\xab\xf5\xf1\x2f\xd3\xc1\x87\xba\xd8\x7e\xab\x2d\xa5\x83\xb9\x49\x77\x47\x9a\x74\xdc\x46\xa4\x0a\x85\x80\x54\x0e\xd7\x68\x86\x3e\x82\x4a\x9f\x20\x7e\x82\xad\xb0\xb9\xa1\x70\xac\xc2\x67\x23\x93\x31\x61\x8d\xc9\xea\xf2\xc4\x21\x39\xf2\xc6\x4e\x6f\xce\xf0\x45\xa2\x9a\x48\xb2\x7b\x90\x1b\x26\x1b\x61\x2f\xa4\x1f\xc1\x2f\xb3\x23\xe4\x4a\x67\x23\x45\x2e\x14\xcf\x39\x17\x22\x40\x91\x00\x46\xcd\x99\x0a\x49\x44\x21\x6d\x49\x7c\x26\xcd\x1a\xf2\x39\xbd\xc9\x6d\xab\x03\x59\x27\x74\x2a\xf5\xa8\x98\xbd\x3d\x3c\x81\xf2\xba\x13\x83\x5c\x14\x32\xed\x01\x1c\x71\x2a\xac\xe7\xdb\xd4\xda\xb7\xa1\x5a\x8b\x17\x00\x79\xb4\xdd\x31\x5d\x35\xee\xa4\x4d\x4e\x1b\x4a\xc1\xcf\xcc\x1a\x33\xe9\x44\x95\xf4\xe5\x68\xbe\xb8\x5f\x7e\xfc\x1e\xc9\x22\x97\x0c\xbf\x6f\xa6\x58\x68\xd5\xee\xee\x4f\x07\x88\x0f\xf7\x9a\xa7\x9c\x0f\x6b\x41\xbc\xdd\x4f\x40\xc7\x0b\x3e\x33\xfb\x54\xbe\x0b\x37\x89\x67\xe8\x24\x91\x4e\x45\x5f\x7e\x45\xa5\xa3\x7c\xb7\x42\xc3\xc1\x2a\x06\xaa\x70\xb3\xa9\x15\x54\xd2\xde\x3c\xb0\xd6\xa1\x3a\x39\x6e\x6c\x7c\xdd\x31\x04\xc1\xa1\x5c\xe6\x2b\xd4\x70\xab\x5a\x25\x71\xf3\xd4\x91\x8c\xce\x0d\xfe\x99\x74\x2a\x23\xdf\x1d\x0f\xfd\x0f\x2f\x71\x8e\x4b\xef\x61\x61\x3e\x37\x42\xb2\xb8\xf6\xee\x75\x0e\xba\x20\xf7\x18\x65\xdf\x7a\x23\xda\xfc\x38\xec\x84\xec\x8f\xdf\x3d\x31\x60\x2f\xbc\x3d\x23\xde\x33\xd9\xa5\x12\xfc\x45\x70\x27\x60\xff\x89\xdb\x39\x91\x9b\x67\x64\xff\x7d\x17\x7b\x54\x0d\x46\x28\x69\x41\xb4\x06\x45\xb5\x0b\xcb\xaf\xae\x20\x74\xc5\xce\xbd\x6d\xcb\x1b\xdf\xef\x13\x51\x56\x1f\x86\xa7\x6b\xe6\xb9\xe9\xf6\xf8\xb4\xbc\x12\xe5\xc4\x0e\x0d\xae\x2f\xf4\xe2\x4f\x11\x6a\x38\x29\xc3\xcd\x75\x7a\x52\x77\x4a\x9e\x51\x54\xc5\x70\xf5\x9d\x25\x1b\x5e\xfd\x46\x51\x4e\xae\xc6\xa7\x0c\xc6\x32\x60\xd9\xe5\x84\xf3\x9e\xc9\x51\x72\xad\x16\x67\x44\xbf\x40\x37\x95\xda\xc5\x32\x23\x98\x58\x52\xa6\x6b\x4c\x3e\x77\xf8\x31\xd6\x29\x91\x85\x55\x0c\xe5\xdf\x04\x61\x28\x00\xf7\x0e\xe7\x34\x6f\xd8\xb5\xc5\x33\xca\x6b\x26\x3b\x62\x29\x17\x6f\xda\x62\x8c\x57\xe1\x04\x5d\xed\xc2\x4d\x20\x97\x2e\x6d\x9b\xce\x55\x7e\xcb\x72\x4a\x04\x4c\x5b\x84\x0d\x1c\xfa\x08\x8f\x4e\x43\x89\x41\xeb\xcf\xa9\xe3\x02\xdd\xa5\xc1\xe4\x13\xcf\xba\x38\x9a\x5c\x10\x4a\x42\x91\xf7\x90\x58\x12\x76\x34\x0d\x26\x71\xfc\x8e\x68\x62\xd1\xa5\x77\xf8\x27\x65\x36\xd0\x4e\x3b\x78\x77\x8c\xdb\x4b\xd3\x85\xcf\xc9\x7a\xd2\xdf\x13\x54\xd2\xf2\xc3\xf6\x2a\xa6\x3c\x3a\x97\xcc\xff\x69\xf3\xfb\x5a\x6e\x96\xf2\xf0\x59\xdd\x05\xca\xd1\xc2\xf5\xf4\xe0\x5d\x81\xdb\xf8\xf5\xff\x5d\xee\xc5\xf3\x0e\x13\xf1\x88\x4a\x69\xf6\xfa\x01\xa8\x71\x5e\xba\x0e\xa8\x75\xdb\xea\x2d\xd7\x4f\x8f\x6d\xd2\x54\x78\xec\x7d\x86\x9a\x02\xe1\x54\x17\x8d\xaf\x8f\x0c\x5e\xea\xe0\x42\x55\xba\x93\xbf\xc5\xa2\xe9\xeb\x52\x11\xa5\x5d\x81\x4a\xfb\x75\x73\xdf\xb3\x15\x07\x81\xe6\x98\x13\x8c\x9f\x76\x88\x24\xa3\x03\xe5\xc4\xd1\xa4\x95\xe0\x08\x61\xf6\xa0\x8d\x48\x93\xfd\xe2\xac\x2e\xed\xd1\x06\xed\xd1\xfe\xec\xbd\x29\x4a\x2b\xf8\xcf\x3b\x60\xa3\xc3\x8d\x3f\xc1\x3e\xa0\x49\x9b\x0e\x59\x82\xa9\xc6\x17\x6e\xa1\x84\x4b\x31\x86\x3f\x8f\xd8\x50\x21\x72\x80\x4f\xff\x98\x6c\x12\x4d\x0e\x27\xdf\xbd\x46\x96\xa2\x5f\x75\xd2\x39\x34\xc5\x04\xed\x2a\x1c\xd0\x89\x20\x74\x30\xd2\x52\xae\xa6\xbc\x16\xf0\xf9\x46\xf6\x3d\xd7\x06\x43\xd7\x76\xac\xae\xf3\x5b\xc9\xf7\x76\x91\x8f\x37\x91\xbf\x4e\x7f\xff\x47\x9d\xe4\x87\x1b\xc4\x1d\x80\x97\xda\xc4\x1d\x30\x0f\x30\x8b\x84\x74\xb1\x65\x38\xbd\x5e\xb7\x78\x76\xf0\xdc\x23\x9f\xd8\xc6\xf0\xf5\xd8\xa7\xe3\xe3\x17\x47\xd8\x2f\x81\xc9\xf0\x5a\x3d\xfd\x01\x56\xfe\x13\x21\xad\x9e\xe9\xba\x3e\x7d\x69\xc3\x40\x55\xa1\xeb\x7a\x09\xf3\xeb\x0c\x37\x00\xe5\xf0\x17\x49\x61\x1f\x76\x0f\x44\x9d\x8d\xa1\xf8\xf5\x24\xdb\x0b\x1a\xcb\x7f\x25\x74\x4a\xec\x91\x70\x22\xbd\xcd\xd7\x14\xcc\xc9\x0a\x23\xf8\xde\x5f\x70\x9c\x92\x5d\x5a\xf9\xf0\x87\x3b\xc3\x50\x36\xd6\x68\x62\xe9\xe1\xf4\xc9\x4d\x32\xdd\x74\x8f\xfa\xe2\xee\xef\x1b\x4e\x17\xc2\x2e\xe3\x43\x6a\x59\x83\x50\xc3\x1b\x16\x72\x99\xd8\xfb\xbb\x02\x7d\x4c\x28\x61\x1a\xf7\xfb\xb7\xf2\x8c\x1b\x84\x5e\x18\x3b\xbe\x34\xa0\xb0\x64\xf0\x5f\x1e\x2d\xf9\x62\x84\xdb\x7b\x88\x40\x33\xa6\xcf\x30\xbc\x2b\x74\x5d\x18\xda\x40\xc6\xfa\x3b\xcf\xb6\xb9\xfd\x99\x1e\xd8\xf3\xfe\x44\xeb\x31\xdd\x50\xd7\xe1\x2d\x85\xaa\xc6\xbf\x0f\x73\xb1\xd8\x85\x8b\x6a\xd9\x0f\xa0\x49\x5a\xf6\x1e\x80\x40\x33\xca\xe5\xf6\xc3\x5d\xce\xd5\x06\xe1\xc7\x16\xe8\x00\xf7\xbf\x01\x00\x00\xff\xff\xb1\xd0\x14\x0d\x57\x3c\x00\x00")

func configYamlBytes() ([]byte, error) {
	return bindataRead(
		_configYaml,
		"config.yaml",
	)
}

func configYaml() (*asset, error) {
	bytes, err := configYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config.yaml", size: 15447, mode: os.FileMode(420), modTime: time.Unix(1485914561, 0)}
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
	"config.yaml": configYaml,
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
	"config.yaml": &bintree{configYaml, map[string]*bintree{}},
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

