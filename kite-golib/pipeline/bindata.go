// Code generated by go-bindata.
// sources:
// templates/feed-errors.html
// templates/root.html
// DO NOT EDIT!

package pipeline

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

var _templatesFeedErrorsHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x53\x4d\x6f\xdb\x30\x0c\xbd\xfb\x57\x10\x5e\xaf\x91\xd6\xb4\xbb\x64\x8a\x2f\xc3\x76\x29\x30\x0c\xeb\xfe\x80\x62\xb1\xb1\x50\x7d\x18\x12\xd3\xc5\x33\xfc\xdf\x07\x59\x72\x92\x36\x40\xab\x93\xf8\xf4\xf8\x48\x3d\x4a\xa2\x23\x6b\x9a\x4a\x74\x28\x55\x53\x01\x00\x08\xd2\x64\xb0\x19\x47\xf6\x1b\x65\xf4\x6e\x9a\x60\x05\xe3\xc8\x7e\x20\xaa\x65\xff\x4b\xf7\xf8\x53\x5a\x9c\xe3\x14\x18\xed\x10\x12\x03\xbe\x87\xe0\x43\x14\x3c\xab\x64\x45\xa3\xdd\x33\x04\x34\xdb\x3a\xd2\x60\x30\x76\x88\x54\x03\x0d\x3d\x6e\x6b\xc2\x23\xf1\x36\xc6\x1a\xba\x80\x4f\xdb\xba\x23\xea\xe3\x86\x73\x2b\x8f\xad\x72\x6c\xe7\x3d\x45\x0a\xb2\x4f\x41\xeb\x2d\x3f\x01\xfc\x8e\xdd\xb1\x2f\x29\xf5\x8c\x31\xab\x1d\x4b\x62\x8d\xe0\xa9\x68\xa9\x3f\x57\xcd\xfb\xb4\x18\x86\xb0\x0a\xfe\x2f\x8c\x27\x28\x2d\x2b\xc3\x5e\xbb\xd5\xce\x13\x79\xbb\x81\xf5\xe7\xfe\xf8\xf5\x15\xa1\x97\x4a\x69\xb7\xdf\xc0\xed\xd5\xd1\x4e\xb6\xcf\xfb\xe0\x0f\x4e\xad\x5a\x6f\x7c\xd8\xc0\x27\xbc\x45\xf9\xb4\x3e\xd3\xa6\xdc\x0b\x2f\xcd\x08\x9e\x2d\xaf\xc4\xce\xab\xa1\xa9\x84\xd2\x2f\xd0\x1a\x19\xe3\xb6\x6e\xbd\x23\xa9\x1d\x86\xba\x5c\xa0\x5b\xa7\x79\x64\xc7\x05\xef\xd6\x4d\x55\xf0\xfb\x84\x67\xcb\xd9\x37\x7f\x70\x34\x4d\x97\x63\xb8\x9e\xdc\x32\x53\xc1\xbb\xfb\xa2\x32\x8e\x41\xba\x3d\xc2\x4d\x84\xcd\x16\x16\xb5\x47\x69\x7b\x83\x71\x2a\x6d\x5f\x74\x97\x9c\x2b\x0e\xd6\x67\x53\x85\x32\x0b\x41\x99\x55\xe7\x83\xfe\x97\x6e\x61\x2e\x28\x99\x46\xcd\x1f\x6d\x31\x92\xb4\xbd\xe0\x8a\xde\x1e\xab\x66\x1c\x6f\x22\x3b\x71\x52\xab\x4a\x5d\x8b\x3c\xfa\x43\x68\x11\x92\x27\xef\xc8\x64\xd6\x62\xdc\x3b\x3a\x0f\x38\x7c\x28\xf3\x80\xc3\x5b\x15\xc1\x95\x29\x36\xce\x61\xeb\x15\xbe\xd6\x98\xf3\x67\x4f\x8b\x95\x39\xed\x4c\x14\x5c\xe9\x97\xa6\x4c\x02\x9d\x9a\xa6\xaa\x40\x82\xe7\xa7\x91\xde\x4a\xfa\xa5\xff\x03\x00\x00\xff\xff\x13\x20\xbb\xf4\xac\x03\x00\x00")

func templatesFeedErrorsHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesFeedErrorsHtml,
		"templates/feed-errors.html",
	)
}

func templatesFeedErrorsHtml() (*asset, error) {
	bytes, err := templatesFeedErrorsHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/feed-errors.html", size: 940, mode: os.FileMode(420), modTime: time.Unix(1565022234, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesRootHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x56\xcb\x6e\xeb\x36\x10\xdd\xfb\x2b\x58\x35\x28\xda\x85\x44\x34\xf7\x76\xe3\x52\x2a\x8a\xe2\x5e\x20\x9b\x26\x70\x8a\xee\x69\x73\x6c\x11\xa1\x48\x81\x1c\xe5\x51\x41\xff\x5e\x90\x94\xac\x47\xe4\xc4\xda\x58\xe4\xcc\x39\x73\x38\x0f\xd1\xac\xc4\x4a\x15\x1b\x56\x02\x17\xc5\x86\x10\x42\x18\x4a\x54\x50\xb4\x6d\xf6\x20\x6b\xf8\x9b\x57\xd0\x75\x24\x25\x7e\xa1\xa4\x06\xf2\x88\x1c\x1b\xc7\x68\x74\x8b\x10\x25\xf5\x13\xb1\xa0\xf2\xc4\xe1\x9b\x02\x57\x02\x60\x42\xf0\xad\x86\x3c\x41\x78\x45\x7a\x70\x2e\x21\xa5\x85\x63\x9e\x94\x88\xb5\xdb\x52\x5a\xf1\xd7\x83\xd0\xd9\xde\x18\x74\x68\x79\xed\x17\x07\x53\xd1\xf3\x06\xfd\x92\x7d\xc9\x7e\xf3\xd0\x71\x2f\xab\xa4\xce\x3c\x59\xc1\xa8\x0f\xda\xc7\x0f\x51\xe3\xbb\x7f\x90\xef\x15\x90\xf6\xbc\xf6\xcf\x8b\x14\x58\x6e\x49\xa3\x1d\x20\xf9\x41\x56\xb5\xb1\xc8\x35\xfe\x7e\x76\xea\x46\xb8\xc8\x1c\x72\x4c\x9f\xb9\x5a\x90\xf8\xb3\xa4\x5c\xc9\x93\xde\x12\x2b\x4f\xe5\x25\xf8\x51\x5a\x87\xa9\x27\x59\x10\xec\xf9\xe1\xe9\x64\x4d\xa3\x45\x7a\x30\xca\xd8\x2d\xf9\x11\x7e\x05\x7e\xbc\x5d\x12\x31\xda\x9f\x89\xd1\x58\x9a\x0d\xdb\x1b\xf1\x56\x6c\x98\x90\xcf\xe4\xa0\xb8\x73\x79\x72\x30\x1a\xb9\xd4\x60\x93\x3e\x0f\xe5\xed\xbc\x6e\x8c\x96\xb7\xc5\x26\xda\x84\x1a\x60\x42\xa5\xa5\xb1\xf2\x3f\x8f\x56\xc9\x98\x36\x26\xb0\xd8\x19\x05\x8c\x0a\x9c\xee\x0a\x4f\xea\x0d\x9e\x50\x88\x39\xe0\x11\xb9\x45\x10\x84\xe3\x2a\x2c\x76\x4b\xd6\x7b\xfd\x89\x6b\x14\xdf\xa5\x96\xae\xfc\x94\x63\x70\x5b\x27\xd9\x35\x9a\xa0\xac\x2e\x88\x6f\xf4\x3f\xb2\xba\xa4\x1f\xd7\x41\xa3\x74\x7c\x87\x6c\x5b\x79\x24\x83\xc7\x37\x6b\xbb\x6e\x56\x66\x4f\xfc\xcd\x5a\x63\xe7\xc4\x23\xb9\x05\x2d\xc0\x06\x97\x39\xcd\x32\x0c\x68\xd1\x0d\x1d\x21\xd4\x50\xcb\xf2\x6b\xf1\x1d\x40\x84\x59\x74\x8c\x96\x5f\x23\x24\xaa\xf2\x96\x60\x98\xeb\x8a\xc1\x4e\x80\x28\xf5\x89\x1c\x3d\xdc\x37\xa8\xdb\x92\x85\x9c\x15\x7c\xdb\x82\x72\x30\x08\x89\xe3\xd5\x37\x53\x58\x4c\x9b\x08\xc7\x4f\x49\x5c\xdb\x45\x02\xb0\x0c\xe2\x19\xc5\xf2\xbd\xe5\x2f\xd3\x68\x5c\x37\x79\x4d\x73\x0b\xa3\x53\x72\x6f\x5b\x84\x8e\x23\x33\x66\xd3\x72\x7d\x02\x72\xe3\xc8\x36\x9f\x9c\x73\x59\xbc\xa5\xe2\xb8\x29\x88\x35\x2f\xae\xe6\x3a\x4f\xda\xf6\xc6\x65\x3b\xf3\xe2\xba\x2e\x29\x98\x43\x6b\xf4\xa9\x08\x9b\x9e\xd3\x17\xb1\xdf\x63\x14\xc5\x3a\x57\x9f\xbd\xf3\x77\x66\xfc\x62\x24\x91\xe9\x4e\x7b\x9e\x4f\xe0\x53\xd4\x9d\x7e\xef\x3e\x4f\xd0\x87\x87\x5b\x08\xea\x55\xdc\x37\xf8\x81\x8c\xe2\xbe\xc1\x6b\x82\x9e\x13\x0f\x3e\xf1\x37\xa1\xd5\x8d\x5d\xe6\xfd\xa2\xbc\x8f\x24\x42\x16\x3a\xe6\x92\xc8\x41\x28\xe3\xfd\xed\x43\x7d\xe3\xa7\x10\xe2\xff\xe1\xdf\xf3\x49\xdd\x7e\xb2\xc0\x9d\xd1\x79\xe0\xdd\x85\x77\x5f\xe1\x38\x17\x3f\xcf\x76\x7f\x61\x94\x5f\x2a\xef\xca\xf9\xc7\x31\x5e\x5b\x33\x3a\xe9\x54\x46\xc3\x48\x0d\x13\x1d\x3d\xcf\x63\xff\xc0\x2d\xaf\x26\x23\xbf\x32\x8c\xf1\xfa\x4b\x1d\x5a\x59\x83\x48\xae\x98\x87\x3a\xcc\x43\xa4\xbe\x72\x18\x7c\xee\xeb\x6c\xb8\x65\x2e\xf5\x47\x70\xfa\x97\xab\x66\xd5\x6b\x9e\xa8\xeb\x92\xc2\xa8\x90\xcf\xfe\xa7\xbf\x0d\x69\xf8\xff\xf2\x7f\x00\x00\x00\xff\xff\xf2\xd9\x78\xa5\xc6\x08\x00\x00")

func templatesRootHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesRootHtml,
		"templates/root.html",
	)
}

func templatesRootHtml() (*asset, error) {
	bytes, err := templatesRootHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/root.html", size: 2246, mode: os.FileMode(420), modTime: time.Unix(1565048438, 0)}
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
	"templates/feed-errors.html": templatesFeedErrorsHtml,
	"templates/root.html": templatesRootHtml,
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
	"templates": &bintree{nil, map[string]*bintree{
		"feed-errors.html": &bintree{templatesFeedErrorsHtml, map[string]*bintree{}},
		"root.html": &bintree{templatesRootHtml, map[string]*bintree{}},
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
