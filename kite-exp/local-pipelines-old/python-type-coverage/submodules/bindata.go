// Code generated by go-bindata.
// sources:
// templates/toplevel.html
// DO NOT EDIT!

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
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
	name string
	size int64
	mode os.FileMode
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

var _templatesToplevelHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xdc\x56\xcd\x72\xd3\x30\x10\xbe\xf3\x14\x8b\xe9\x91\x4a\x34\x85\x4b\x50\xcc\x30\x3d\x71\x28\x17\x9e\x40\x96\x44\xed\xe0\x3f\xa4\x0d\x93\x8c\xc9\xbb\x23\xc9\xf1\x4f\x9c\x38\x71\x32\xfc\x4d\x73\x68\xbd\xbb\xde\xdd\xef\xdb\x6f\xb7\x0d\x7b\x29\x0b\x81\x9b\x52\x41\x8c\x59\x1a\xbe\x60\xf5\x2f\xb0\x1f\x16\x2b\x2e\xeb\x47\x6b\x60\x82\xa9\x0a\x1f\x8a\x55\x8e\x86\xd1\xda\x6a\x62\x46\xe8\xa4\x44\x70\x65\x16\x01\xaa\x35\xd2\x25\xff\xc1\x6b\x6f\x00\x46\x8b\x45\x10\x23\x96\x66\x4e\xa9\x28\xa4\x22\xcb\xef\x2b\xa5\x37\x44\x14\x19\xad\x1f\x6f\x67\xe4\x8e\xbc\x25\x59\x92\x93\xa5\x09\x42\x46\xeb\xdc\x6b\x1a\x64\x7c\x2d\x64\x4e\xa2\xa2\x40\x83\x9a\x97\xce\x70\x8d\x5a\x07\xbd\x27\xf7\xe4\x1d\x5d\x9a\xce\x35\xde\x38\x4d\xf2\x6f\xa0\x55\xba\x08\x0c\x6e\x52\x65\x62\xa5\x6c\xc3\x1e\x0e\x61\x4c\x00\xb1\x56\x5f\x2f\x47\x60\x53\x07\x10\x5c\x31\x8b\xc1\x35\xed\xa8\xbb\xbe\x8d\x05\x10\x15\x72\x03\x55\x6b\x02\x94\x5c\xca\x24\x7f\x9a\xc3\xec\x4d\xb9\x7e\xdf\x06\xb6\x4d\x3e\xed\x15\x60\xb4\x93\x94\xb9\x4a\x6d\x97\xf8\x2e\xfc\xa2\x04\x26\x45\x6e\xc5\xb5\x46\xab\x3a\x8f\x52\x05\x22\xe5\xc6\x58\xc2\xde\xf0\x3f\x6f\x2d\xe8\xa4\x54\x32\xe8\x90\x31\xd4\x21\x43\x19\x32\xbe\x9b\xc7\xab\xc7\x42\xae\xec\xcc\x82\xf0\x31\x31\xc6\x62\x84\x4f\x59\x59\x68\x54\x12\x76\x11\x46\xb9\xa5\xeb\x72\xa8\x4d\x3e\x5b\xe9\x23\xa2\xee\x55\xab\x9d\xe0\xbc\x49\xb4\xc2\x0b\xca\x3d\x38\x3e\x83\x6a\xde\x77\xae\x98\x7d\x76\xec\xbb\xb1\xcd\x20\x91\x8b\x60\x02\xd1\x78\x76\xdd\x4c\xfb\x37\xd8\x90\xe9\x99\xfe\x95\xb0\xee\x62\xc1\xc5\x87\x31\x7f\xb1\xc3\xd0\x60\x40\x74\xd0\x86\x61\x7f\x39\xdc\xa7\xaa\x34\xcf\x9f\x14\xdc\x70\x01\xf3\x05\x90\x1d\xd1\x1d\xbd\xed\x76\xd0\x76\x80\xd1\xb9\x64\x58\x55\x36\x9b\xb8\x09\x6f\xb7\x7e\xb2\xa3\xef\x78\xcc\xf0\x13\x4a\xad\x10\x37\xde\xaa\x53\x3e\x0c\xf8\xd1\xfd\x4e\x55\xa5\x72\xd9\x03\x63\xe3\x7b\x5b\x7e\x4a\xbd\xf3\xcb\xf5\xcf\x24\x6c\x41\xfc\x15\x85\xf3\x23\x0a\xfb\xe1\x4c\x53\x99\x45\x4e\xc4\x5c\x90\xcf\x3c\x53\x4e\xb5\x28\x1c\x13\x7b\xd4\xdf\x94\x68\x95\xdf\xd5\x00\x38\x29\xff\xc1\x9a\xba\x1a\x1e\xba\xdd\x25\xe3\x0f\xb2\xfe\xff\x35\x60\x72\x94\xcb\x38\xc4\x69\xdb\x3c\x69\x9f\x0f\xd8\x1f\x61\xb4\xbf\xd2\x23\x68\x4f\x8e\x73\xb2\xff\xf7\xdc\xd3\xb4\xbf\xae\x7f\xf4\x9a\x7c\xbb\xff\xe1\x98\x3c\xa9\xd7\x37\xdc\x2f\x61\xef\xac\xba\x19\x5d\x72\x55\xbe\xda\x95\x37\x75\xd5\xfd\xf0\xe7\x76\x3c\x97\x6c\xb4\x1d\xb3\x0f\xd9\x5d\xf5\xdf\x89\x7f\x05\x00\x00\xff\xff\xfd\x49\x75\xed\x2b\x0b\x00\x00")

func templatesToplevelHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesToplevelHtml,
		"templates/toplevel.html",
	)
}

func templatesToplevelHtml() (*asset, error) {
	bytes, err := templatesToplevelHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/toplevel.html", size: 2859, mode: os.FileMode(420), modTime: time.Unix(1478031646, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"templates/toplevel.html": templatesToplevelHtml,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"templates": &bintree{nil, map[string]*bintree{
		"toplevel.html": &bintree{templatesToplevelHtml, map[string]*bintree{
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
