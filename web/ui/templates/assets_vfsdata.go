// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package templates

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Assets contains the project's static assets.
var Assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2018, 7, 27, 12, 17, 4, 100006584, time.UTC),
		},
		"/_base.html": &vfsgen۰CompressedFileInfo{
			name:             "_base.html",
			modTime:          time.Date(2018, 7, 27, 8, 8, 39, 327241265, time.UTC),
			uncompressedSize: 3007,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x56\xdf\x6f\xdc\x36\x0c\x7e\xef\x5f\xc1\xa9\xc5\xda\x02\xf3\x19\x43\x5f\x86\xd6\xf6\xd0\x26\xe9\x1a\xa0\x58\x0f\xc9\xad\xd8\x30\x0c\x01\xcf\xa6\x6d\xa5\xb2\xe4\x4a\xb4\x9b\xc3\xe1\xfe\xf7\x41\xe7\x1f\xf3\x39\xb9\x64\xeb\x86\x3d\x59\xa6\x3f\x7e\xa4\xc8\x8f\xb2\xa2\x6f\x4e\x3f\x9c\xac\x7e\x5b\x9e\x41\xc9\x95\x4a\x1e\x45\xfe\x01\x0a\x75\x11\x0b\xd2\x22\x79\x04\x10\x95\x84\x99\x5f\x00\x44\x15\x31\x42\xc9\x5c\x07\xf4\xb9\x91\x6d\x2c\x4e\x8c\x66\xd2\x1c\xac\x36\x35\x09\x48\xbb\xb7\x58\x30\xdd\x70\xe8\xa9\x5e\x41\x5a\xa2\x75\xc4\x71\xc3\x79\xf0\x83\x98\xf2\x68\xac\x28\x16\xd6\xac\x0d\xbb\x89\xaf\x36\x52\x67\x74\xf3\x9d\x36\xb9\x51\xca\x7c\x19\x7c\x58\xb2\xa2\x64\x69\x4d\x45\x5c\x52\xe3\x60\x25\x2b\x82\x4b\xb2\x92\x1c\x9c\x18\xa5\x28\x65\x69\x34\xa0\xce\x60\x69\x4d\x4a\xce\x49\x5d\x78\x40\x4b\x36\x0a\x3b\xf7\x8e\x4a\x49\xfd\x09\x2c\xa9\x58\xb8\xd2\x58\x4e\x1b\x06\x99\x1a\x2d\xa0\xb4\x94\xc7\x62\xbb\x85\x1a\xb9\x5c\x5a\xca\xe5\x0d\xec\x76\xa1\x63\x64\x99\x86\xb2\x2a\xc2\x1c\x5b\x0f\x5d\xc8\xd4\xfc\xd8\xc6\xdb\x2d\xac\x1b\xa9\xb2\x8f\x64\x9d\x8f\xbd\xdb\x0d\xd9\xba\xd4\xca\x9a\xc1\xd9\xf4\x38\x5f\x4b\x3a\x33\x36\xbc\x76\xe1\xf5\xe7\x86\xec\x66\x51\x49\xbd\xb8\x76\x47\x78\xa3\xb0\xe3\xfc\xe7\x01\xd6\xc6\xb0\x63\x8b\x75\xf0\x62\xf1\x62\xf1\xbd\x0f\x38\x9a\xfe\x6e\xcc\x49\xe1\x78\x53\x53\xdf\xe2\xd4\x39\xd1\x17\x92\x37\x8a\x5c\x49\xc4\x0f\x55\xf1\x48\x52\xa9\x9b\x67\x95\xba\x63\x69\xfd\x77\xc9\xf8\xa8\xf5\x28\xa9\xfb\x42\x4e\xab\xde\x25\x00\xd0\xa2\x85\xe5\xeb\xd5\xbb\xab\xe5\xc5\xd9\xdb\xf3\x5f\x21\x86\x5b\x81\xc4\xab\x09\xf6\xcd\x2f\xe7\xef\x4f\xaf\x3e\x9e\x5d\x5c\x9e\x7f\xf8\xb9\x47\xcf\x23\x0d\xf8\x27\xcf\xf2\x46\x77\x8a\x7e\xf6\x1c\xb6\xbd\xd5\xdb\x9f\xfe\x9e\x21\x63\xc0\xa6\x28\x94\xdf\xbb\x31\x8a\x65\x2d\xfe\x78\xfa\x7c\xd1\xaf\x9f\x3d\xef\xe1\xbb\x6e\x31\x6b\xe3\x76\xcb\x54\xd5\x0a\x99\x40\xf8\xe1\x16\xb0\xd8\xed\xfc\xa4\x87\xdd\xa8\xfb\xe5\xda\x64\x9b\xbe\xce\x1a\x5b\x48\x15\x3a\x17\x0b\x8d\xed\x1a\x2d\x74\x8f\x40\xea\x96\xac\xa3\xe1\x35\x97\x37\x94\x05\x6c\x6a\x31\xd4\x27\xca\xe4\xe8\xea\xe7\x1b\xa5\x26\x1b\xe4\xaa\x91\xd9\x88\x39\x44\xf5\x54\x3e\x0f\xb2\x13\x8c\xcf\xa8\x61\x36\xba\x6f\x78\xf7\x22\x66\x6e\x5d\x49\x20\x35\x4a\x61\xed\x28\x13\x70\x50\xa9\xc1\x3e\x98\xd1\x16\xc4\xb1\x78\xdc\x79\x0b\x40\x2b\x31\xa0\x9b\x1a\x75\x46\x59\x2c\x72\x54\x1e\xbb\xb7\xfa\xec\xad\x51\x63\xa8\x83\xd4\xbc\x2e\x6a\xd4\x43\x32\xce\x06\x46\xab\x8d\x48\x56\x5d\x3a\x1a\x5b\x59\xa0\xef\x64\x14\x7a\xdc\x3d\xae\xfe\x68\x09\xf6\xf4\xff\x17\x34\x0a\xbb\x52\x1e\xd8\x70\x56\xd7\xb5\x45\x9d\x1d\x1d\x25\x31\x39\x94\xa3\x10\x27\x8d\x0d\x33\xd9\xce\xfa\x2c\xb3\xb1\x84\xb3\x20\x43\x77\xc6\xf6\x1d\xb6\xbf\x51\x13\xfc\x20\xb9\xc9\x52\x51\xce\xb3\xae\x6c\xb7\x4f\x52\xa3\x9d\x51\xe4\xe0\x65\x0c\xc3\x7a\x89\x5c\xee\xf5\x3e\x45\xca\x1c\x46\xf0\xec\x63\xa4\x64\x12\xe1\xb8\xfb\x09\x4c\x24\x27\xfd\xda\xef\x3b\x0a\x95\x9c\x27\x00\xa4\x33\xb8\x9f\x6f\x56\x4d\x54\x64\xd9\x89\xe4\xf5\xfe\x79\x37\xef\xfd\x0c\x85\xc5\xba\x14\xc9\x4f\xfe\x71\xd4\x7f\x28\x66\x66\x4d\x9d\x99\x2f\x7a\x56\xba\xbd\x08\x3a\xfe\xc7\x62\x8e\xed\x07\x6a\x36\x5d\x23\x13\x58\xa3\x26\x23\xba\x9f\x9f\x12\x5d\x6d\xea\xa6\x8e\x05\xdb\x86\x8e\x8c\x5a\x72\xc9\xc8\x8d\x3b\x14\x6f\x8a\x96\x78\x54\xee\x81\xbe\x6e\x29\x63\x4c\xb0\x22\xdd\xdc\xda\xd1\x43\x75\x73\xfb\xe8\x22\xb9\x68\x34\xfb\xab\xc5\xb7\x58\xd5\xaf\xe0\x8d\x3f\x9f\xe1\x5c\xe7\xc6\x56\xfd\x10\xdf\x55\xd2\x87\xe9\x73\x85\x85\xf3\x8a\xa9\x2a\xd4\x59\xf0\x5e\x6a\x82\xb7\xde\xf6\xb5\x84\xa9\xd1\xb9\x2c\xf6\x1a\xcc\x65\xd1\xd8\x7f\x95\x9d\x6d\x14\xed\xf7\x7e\x54\xcc\x0f\x73\x74\x07\xaa\x13\xc9\xaa\x5b\x7c\x2d\x8f\x23\xdb\xca\x94\x82\x4c\xba\xd4\xb4\x64\x37\x22\xb9\xec\x4c\x70\x3a\x98\x8e\x71\x47\x61\xa3\x66\x62\xbf\x73\x7c\x8e\xa9\xdd\x5f\x6e\xdd\xcb\x70\x7a\x29\x90\x26\xcc\x4c\xea\x04\x0c\x3f\x8c\xab\xb5\x42\xfd\x49\x24\xef\x48\xd5\xb7\x04\x39\x0f\x77\x98\xd0\xc1\x91\x38\x79\x89\x42\x8d\xed\x1d\xbf\xe7\xfe\x56\xfc\xd7\x1f\xba\xfb\x2f\x47\x61\x77\x5d\xff\x33\x00\x00\xff\xff\xeb\x41\x6b\xee\xbf\x0b\x00\x00"),
		},
		"/alerts.html": &vfsgen۰CompressedFileInfo{
			name:             "alerts.html",
			modTime:          time.Date(2018, 6, 25, 7, 35, 55, 128143670, time.UTC),
			uncompressedSize: 2536,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x56\xcd\x6e\x1b\x37\x10\xbe\xfb\x29\x06\x8b\x1c\x5a\xa0\x12\x11\xa0\xa7\x80\xda\xc2\xc8\xa5\x87\x24\x28\x62\xd7\x57\x83\x22\x47\x5e\x26\x34\xb9\x20\x29\xd9\x02\xcb\x77\x2f\x48\xee\xae\x56\xfb\xd3\xb8\x40\x2e\xc2\x92\x33\xf3\xcd\x37\xbf\x54\x08\x02\x0f\x52\x23\x54\x0d\x32\x51\xc5\x78\x03\x40\x95\xd4\xdf\xc1\x9f\x5b\xdc\x55\x1e\x5f\x3d\xe1\xce\x55\x60\x51\xed\x2a\xe7\xcf\x0a\x5d\x83\xe8\x2b\x68\x2c\x1e\x76\x55\x08\xd0\x32\xdf\xfc\x65\xf1\x20\x5f\x21\x46\xe2\x3c\xf3\x92\x27\x1b\xc2\x14\x5a\xef\xb6\xdc\xb9\x3f\x4e\xbb\x10\x60\x7f\x94\x4a\x3c\xa0\x75\xd2\x68\x88\xb1\xaa\x93\x33\xc7\xad\x6c\x3d\x38\xcb\xd7\xc1\xbe\x0d\x58\xdf\xd6\xa0\x28\x29\x40\xf5\x4d\x08\xa8\x45\x8c\x37\x37\x97\xd8\xb8\xd1\x1e\xb5\x4f\xe1\x51\x21\x4f\xc0\x15\x73\x6e\x97\xaf\x99\xd4\x68\x37\x07\x75\x94\xa2\xf0\x69\xde\xd7\xb7\xd9\x17\x25\xcd\xfb\x7c\x33\xb2\x70\x8d\x79\xd9\x30\xad\x4d\xe2\x65\xb4\xcb\x26\x00\x40\x65\xaf\xf1\xa4\xce\x6d\x23\xb9\xd1\x30\x7c\x6d\x8e\x9a\x37\xc8\xbf\xa3\x48\x34\x65\x67\x02\x40\xf7\x47\xef\x8d\xee\x32\x5d\x0e\xd5\xaa\x27\xf0\xd2\x2b\x2c\x02\xb8\xa2\x70\x37\xb9\xa1\xa4\x60\x65\xf2\x44\xc8\x53\xfe\xf0\x6c\xaf\xb0\x47\x2f\x87\xfc\xbb\xd9\x1b\x2b\xd0\xa2\xe8\x8e\xdc\x28\xc5\x5a\x87\xa2\x8b\x8d\xfa\xbd\x11\xe7\xf2\x1d\xc2\xbb\x5c\x87\x3b\xcf\x3c\xde\x9b\xaf\xe6\xe5\x63\xc2\x83\x0f\x3b\xd8\xde\x2e\x08\x72\x3b\x25\x33\xcb\xf4\x13\x76\x3a\x52\x3f\x7d\x3d\x2a\xec\x85\x05\x95\x7b\x79\xc2\x92\xf7\x82\x36\xba\x18\x14\xa9\xb7\x7d\x00\x21\x48\x2d\xf0\x15\x96\xf9\x6c\xf3\x45\x8c\x90\xa5\x8f\xa9\xb5\xd1\x56\x43\xe2\x81\x7a\x51\x5f\x4a\x96\x6b\xc4\x1b\x3c\x59\xa3\x37\xc2\xbc\xe8\x52\x26\xa0\xfb\x3a\x84\xed\x17\xf6\x8c\x31\x52\xb2\xaf\xe1\x97\x10\x14\x6a\xb8\x62\x9b\x9c\xe4\xe3\xaf\x94\x78\xd1\xbb\xa0\xc4\xdb\x7a\xce\xba\xd0\x11\xe8\x99\x54\x6e\xc2\x67\x38\x94\x8e\x1b\x9f\x01\x68\x6b\xb1\xa6\xdc\x08\x4c\x94\xfe\xbc\xff\xfc\xe9\x4e\xcb\xb6\x45\x3f\x9a\x97\x44\x32\x6b\x50\x92\xb4\xc7\x78\x64\x02\x18\x82\x3c\x4c\xc3\x18\xeb\xbf\xb5\x57\x1a\x73\x42\x3b\xf4\x8d\x16\xa8\x1d\x8a\x2e\xe9\xa8\xf0\x19\xb5\x77\x8f\x59\x5c\x4d\xe2\xb9\xe4\x64\x22\x49\xb2\xa6\xfe\xc4\xf6\xa8\x1c\x25\xbe\x59\x92\xe6\xea\xae\x09\x4b\xe7\xc0\x9d\xd4\x7c\x55\xe7\x81\xa9\xe3\x82\x70\x5c\xb5\x3e\x51\xa5\x73\xd7\x73\x95\x63\x99\xfb\x10\xd3\xab\x11\x96\x4a\xc1\xfd\x06\xef\x4e\x89\x45\xee\xf6\x12\xee\xf6\x33\x6b\x27\xd8\x1d\x9c\x6b\x99\xee\xf3\x95\xad\x21\xff\x6e\x5a\x2b\x9f\x99\x3d\x57\x75\x08\x05\x35\xc6\x34\x1a\x05\x39\xc6\x8a\x92\x64\xb9\x44\xa5\xac\xc8\x89\x1b\x32\xa7\x9d\x27\x65\xec\x3e\x17\xb7\x94\x78\x13\x42\x37\x69\xf0\x0f\x8c\xe7\xb0\x0c\x61\x8c\x90\xd6\x37\x3e\x4a\x2d\x24\x67\xde\x58\x48\xaf\xc9\xe6\xd8\xb6\x68\x39\x73\x98\x68\xf7\x93\xda\x31\x5d\xa3\x10\x42\xbf\x11\xfc\xf6\xef\xfb\x8f\x49\x7f\x55\xf1\xa1\x04\x3f\xd7\x58\x2a\x2f\xc8\x03\x6c\x6f\x2f\xbb\x73\xa1\x06\xa9\x57\xf3\xbb\xb7\xab\x84\x74\xad\x62\xe7\x0f\xda\x68\xac\xae\x87\x7a\xe1\x51\x18\x21\x34\xc0\x8d\x4a\x11\xee\xaa\xdf\xab\xfa\x76\xbc\xab\x7f\xdc\x84\x3f\x83\x80\xb8\x22\x30\x6b\x08\x2a\xd4\xff\x6a\xd8\xff\xce\x58\x8f\xe9\x47\x7d\x49\x89\xf0\x73\x17\x49\x2b\x15\xad\xef\x58\x4a\xc4\xe2\xe4\x2c\xb5\x6b\xde\x6b\x33\xda\x6f\x2b\xfb\x1c\x6f\x7e\x47\x49\xde\x5c\xd7\x7b\xf3\x5a\x69\x79\xe5\x87\x80\xca\xe1\xf8\xd5\x5a\xdd\xf4\x00\x5f\x4c\x19\x1d\xa9\x9f\xc0\xa6\x67\x11\xca\x9f\x16\xf1\x63\x27\x03\x15\x4a\x86\x37\x7a\x20\xdd\xed\xfc\x5e\xed\xdf\x00\x00\x00\xff\xff\x1b\x6e\x06\x91\xe8\x09\x00\x00"),
		},
		"/config.html": &vfsgen۰CompressedFileInfo{
			name:             "config.html",
			modTime:          time.Date(2018, 2, 12, 15, 32, 2, 307685262, time.UTC),
			uncompressedSize: 175,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8e\xc1\x0a\xc2\x30\x10\x44\xef\xf9\x8a\x35\xf7\x58\xe8\x39\xe6\xe2\x97\x84\xee\xc6\x2c\x94\xad\xa4\x69\x11\x96\xfd\x77\x69\x51\xf4\x36\xc3\x7b\x30\xa3\x8a\x54\x58\x08\x7c\xa5\x8c\xde\x2c\x5e\x42\x00\xe1\x17\x84\x90\x54\x49\xd0\xcc\xb9\x9f\x35\x2d\xd2\x49\xba\x37\x73\x00\x11\x79\x87\x69\xce\xeb\x7a\x3b\x41\x66\xa1\x16\xca\xbc\x31\xfa\xe4\x00\x00\x62\x1d\x81\xf1\xa4\x85\x1f\x5b\xcb\x9d\x17\xf1\xe9\xfe\x5f\xe3\x50\xc7\x8f\xfd\x6c\x94\x54\xaf\x66\x71\x38\xe2\x31\x31\x20\xef\xc9\x7d\x9f\xbc\x03\x00\x00\xff\xff\x41\xfa\xfc\xb0\xaf\x00\x00\x00"),
		},
		"/flags.html": &vfsgen۰CompressedFileInfo{
			name:             "flags.html",
			modTime:          time.Date(2018, 2, 12, 15, 32, 2, 307685262, time.UTC),
			uncompressedSize: 433,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x90\xcd\x4e\xc4\x20\x10\xc7\xef\x7d\x8a\x91\xec\x51\x6c\xb2\x47\x43\x7b\x31\xf1\xe4\x4b\xd0\xce\x74\x21\x76\xa1\x01\x5a\xdd\x10\xde\xdd\x80\xa5\xd5\x0b\xc9\x8f\xff\x47\x66\x26\x46\xa4\x49\x1b\x02\xa6\x48\x22\x4b\x49\x3c\x71\x0e\x46\x7f\x03\xe7\x7d\x8c\x64\x30\xa5\xa6\x39\x5d\xa3\x35\x81\x4c\x60\x29\x35\x00\x02\xf5\x06\xe3\x2c\xbd\xef\x8a\x20\xb5\x21\xc7\xa7\x79\xd5\xc8\xfa\x06\x00\x40\xa8\x2b\x68\xec\x98\x0f\xd2\x85\x75\x99\x66\x79\xf3\xac\x7f\xb3\xf7\xbb\x34\xc8\x3f\x72\xe5\x7b\xfe\x13\xad\xba\xee\x89\x20\x87\x99\x6a\xeb\x2f\x94\x97\x8f\xd6\x20\x19\x4f\xb8\xf3\x60\x1d\x92\x3b\xd0\x07\xa7\x97\x83\x94\xdd\xc8\xed\x43\xe4\xd2\xc1\xe2\xa3\x12\x40\x8c\x4e\x9a\x1b\xc1\xe5\x93\x1e\xcf\x70\xd9\xe4\xbc\x12\xbc\x76\xf0\x02\x65\xaf\x1a\x72\x67\x22\xa3\x02\x3f\xda\x85\x3a\xe6\xec\x17\xeb\x63\xcc\xe9\x94\x44\x1b\xd4\x7f\x1f\x66\xad\x74\x16\x15\x4f\x55\xb4\x7f\x3b\xeb\x79\x0f\xed\x1c\x52\xb4\x65\x8d\x0c\xa2\x45\xbd\xf5\x4d\x35\xff\x04\x00\x00\xff\xff\xea\x90\xd3\xc6\xb1\x01\x00\x00"),
		},
		"/graph.html": &vfsgen۰CompressedFileInfo{
			name:             "graph.html",
			modTime:          time.Date(2018, 4, 24, 13, 51, 38, 345724084, time.UTC),
			uncompressedSize: 2296,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x55\xbb\x92\xdb\x3a\x0c\xed\xfd\x15\x1c\xf6\xb2\x8a\x6d\x2d\xdf\xb9\x45\x26\x6d\xaa\xb4\x3b\x14\x09\x2f\x61\x53\xa4\x42\x80\x5a\x6b\x35\xfa\xf7\x8c\x1e\x56\x1c\x4f\x9c\xec\x4e\xac\x34\x36\x48\x1d\xe0\x1c\x3c\x24\x74\x9d\x81\x03\x7a\x10\xd2\x82\x32\xb2\xef\x37\x42\x08\xb1\x73\xe8\x4f\x82\xdb\x1a\x0a\xc9\x70\xe6\x5c\x13\x49\x11\xc1\x15\x92\xb8\x75\x40\x16\x80\xa5\xb0\x11\x0e\x85\xec\x3a\x51\x2b\xb6\x5f\x22\x1c\xf0\x2c\xfa\x3e\x27\x56\x8c\x7a\xf0\xc9\x5f\xa2\xaa\xed\x56\x13\xfd\xd7\x14\x5d\x27\xca\x84\xce\x7c\x85\x48\x18\xbc\xe8\x7b\xb9\xdf\x3c\x8e\xae\x01\x6f\x42\xcc\x23\xea\x13\x59\xf5\xba\x18\xdb\x0a\xfd\xef\x14\x3c\x5a\x00\x04\xaf\xc8\x28\x9f\x95\x21\x30\x71\x54\x75\x66\x14\x03\x63\x05\x35\xea\x13\xc4\xfc\xde\x83\x3f\x29\x9d\xa4\x92\x8e\x58\xb3\xa0\xa8\xdf\x5f\x8b\xf9\x6c\x9e\xb6\xcd\xd3\xf6\x78\x8f\x60\x97\x4f\xb1\xf7\x8f\x20\x72\xaa\x0d\x89\xc7\x94\xd6\x24\xfc\xa9\xcb\x2b\x10\x55\xa1\x02\xcf\xf3\xdf\x3f\x21\xc9\x86\x81\x78\x0b\x1e\xb2\x57\x64\x3b\x8c\x88\x5a\x8b\xf7\x2f\x47\x75\x05\x45\x0b\xdf\x53\x36\xbc\x8d\x6a\xf8\x28\xfd\xf2\x72\x2d\x01\x87\xf4\xf6\xd6\x4e\xbf\xef\x09\xff\xf1\x56\x27\x62\xa5\x2d\x2c\xc6\x5a\x89\x1c\x29\x3f\x7e\x4b\x10\xdb\x2d\x81\x03\xcd\x18\x56\xa6\xb1\x81\x4f\xd0\xd2\x63\xab\x76\x9c\x57\x48\x8e\xde\xc0\xf9\xc3\xb1\xd1\x14\x72\xf4\x7f\x66\xa8\x6a\xa7\x18\xe4\xf5\x57\xfe\x9c\x59\xe5\x8d\x83\x52\x45\xca\x16\xc4\x55\xb0\xae\x03\x6f\xfa\x7e\xb3\xf9\xb1\x26\x75\xf0\x0c\x9e\x97\x4d\x69\xb0\xb9\xa2\x19\x9e\x2a\xf4\x10\xa5\xd0\x4e\x11\x15\x72\xb9\xc9\x0e\x2e\xa1\x99\x17\xce\xe2\x3a\xa3\xc6\x12\x66\x16\x89\x43\x6c\xaf\x30\x23\x0e\x2f\xa8\x17\xd7\xd6\x16\x75\xf0\x62\xb1\xb2\xe4\xb5\x05\x7d\x02\x33\xe8\xc6\x1b\xcf\x32\x31\x07\x3f\xa7\x3c\x1d\x16\x61\x04\x2a\x6a\xbb\x70\x0a\x46\x76\x70\xb9\x16\x75\x84\x06\x43\x22\x31\x28\x43\x20\xb9\xff\xe4\x55\xe9\x60\x3c\xb7\x62\xf6\xda\xe5\x53\xd0\xab\xa4\x72\x83\xcd\x3c\x46\x93\xb9\xb9\xcd\xf5\x5e\x45\x06\xcc\x7e\x87\xbe\x4e\x7c\x81\x96\xec\x45\xc9\x3e\xab\x23\x56\x6a\x14\x39\x66\x42\xa9\xac\x90\xa5\x68\x94\x4b\x50\xc8\xff\x8d\x11\x9f\x87\xf2\xcb\xb1\x13\xca\x98\xe7\xb1\x1b\x43\x45\x6e\xd5\x5c\x3a\xfa\x3d\x00\x00\xff\xff\xc8\x49\xb1\xcd\xf8\x08\x00\x00"),
		},
		"/rules.html": &vfsgen۰CompressedFileInfo{
			name:             "rules.html",
			modTime:          time.Date(2018, 7, 27, 8, 8, 39, 327241265, time.UTC),
			uncompressedSize: 1105,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x53\xc1\x6e\x13\x31\x10\xbd\xe7\x2b\x46\x86\x03\x1c\x36\x86\xde\x40\xde\x45\x95\x40\xe5\x00\x15\x6a\x2b\x0e\x20\x40\xce\x7a\xb6\x6b\xe1\x78\x57\xf6\x6c\x68\x6a\xf9\xdf\x91\x9d\xa4\xd9\x0d\x89\x2a\x9a\x43\x94\x99\xe8\x79\xde\xbc\x37\x2f\x04\x85\x8d\xb6\x08\xac\x45\xa9\x58\x8c\x33\x61\xb4\xfd\x0d\xb4\xee\xb1\x64\x84\x77\xc4\x6b\xef\x19\x38\x34\x25\xf3\xb4\x36\xe8\x5b\x44\x62\xd0\x3a\x6c\x4a\x16\x02\xf4\x92\xda\x2f\x0e\x1b\x7d\x07\x31\x72\x4f\x92\x74\x9d\x30\xdc\x0d\x06\xfd\xbc\xf6\xfe\xdd\xaa\x0c\x01\x16\x83\x36\xea\x2b\x3a\xaf\x3b\x0b\x31\xb2\x6a\x16\x02\x5a\x15\xe3\x6c\xb6\x27\x51\x77\x96\xd0\x52\xe2\x01\x20\x94\x5e\x41\x6d\xa4\xf7\x65\xfe\x43\x6a\x8b\xae\x68\xcc\xa0\x15\xab\x66\x00\x00\xa2\x3d\xab\xae\xd2\x18\xc1\xdb\xb3\x6d\x8b\xe4\xc2\xe0\x0e\xb6\x29\xf2\x77\xb1\xe8\x9c\x42\x87\x3b\x2c\x40\x08\x4e\xda\x5b\x84\x79\x7a\xe2\xc2\x75\x43\xef\xf3\xdc\xcd\x47\x50\x52\xa4\x7a\xa8\x53\xc7\x8d\xcb\xd4\x50\x55\xa2\x20\xe4\x56\x8e\x67\x21\x38\xbc\xc2\xde\xc8\x1a\xcf\x8d\x01\xf6\xe2\xfb\x4f\x59\xdc\x9f\x17\xdf\x5e\x15\x6f\x7e\xbc\x64\xc0\x9e\xbf\x66\x30\xbf\x94\x4b\x8c\x91\x81\x95\x4b\x4c\x1a\xfe\x07\xa6\x0a\x61\xfb\x33\xaf\x2c\x38\xa9\xe3\x9c\x42\x68\x87\xa5\xb4\xfa\x1e\xdf\x0f\x4e\x52\x12\x7d\x7e\x81\xf4\x61\x25\xcd\x90\xcb\x5d\x7b\x7e\x8d\x75\x67\x95\x3f\xf1\xa2\xe0\xe3\xad\x05\x3f\x50\x45\xd0\xa2\x53\xeb\xc7\x54\x82\x7c\x39\x25\x6b\x3a\x4b\xc5\x1f\xd4\xb7\x2d\xbd\x5d\x74\x46\xb1\xec\xde\xd1\x25\x4e\x43\xf6\x2b\xc0\x8d\x5e\xfe\x83\x9e\x12\x3e\x70\x79\x6c\xf0\x09\xaa\xdb\xcb\x49\xd7\xfb\xab\x46\x63\xb2\xe4\x1f\x6f\x3e\x7f\xba\xb6\xba\xef\x91\x46\x07\x9f\x24\x3b\xa6\xff\x93\xb4\x7f\x6c\x8b\x4d\x56\x1e\x2a\xe3\x71\x72\xac\x6e\xea\xc1\x01\xab\xcb\x0e\x72\x1c\x61\x13\x34\x35\x9d\xa4\x26\x06\xbb\x7d\x3e\xc6\x33\x05\x1f\x59\x2d\x78\xce\x54\x2a\x04\x57\x7a\xb5\x0f\xf3\xdf\x00\x00\x00\xff\xff\x8b\x9d\x45\x0e\x51\x04\x00\x00"),
		},
		"/service-discovery.html": &vfsgen۰CompressedFileInfo{
			name:             "service-discovery.html",
			modTime:          time.Date(2018, 6, 25, 7, 35, 55, 128143670, time.UTC),
			uncompressedSize: 2871,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x56\x51\x8f\xe3\x34\x10\x7e\xcf\xaf\x18\x99\x3e\x00\xda\x36\xb7\x27\xc1\x43\x71\x83\x90\xee\x05\x09\x24\x24\x4e\xbc\x20\x84\x9c\x78\xda\xcc\x9e\x6b\x47\xb6\x9b\x6b\x65\xe5\xbf\x23\xc7\x49\x9a\x94\xf6\x80\x85\x47\x56\xda\x95\x3d\x33\xfe\xe6\x9b\xc9\x37\xf6\x86\x20\x71\x4f\x1a\x81\xd5\x28\x24\xeb\xba\x8c\x2b\xd2\x1f\xc0\x5f\x1a\xdc\x31\x8f\x67\x9f\x57\xce\x31\xb0\xa8\x76\xcc\xf9\x8b\x42\x57\x23\x7a\x06\xb5\xc5\xfd\x8e\x85\x00\x8d\xf0\xf5\x4f\x16\xf7\x74\x86\xae\xcb\x9d\x17\x9e\xaa\x78\x26\xf7\xc2\x1e\xd0\xbb\x4d\xe5\xdc\xb7\xed\x2e\x04\x28\x4f\xa4\xe4\x2f\x68\x1d\x19\x0d\x5d\xc7\x8a\x8c\xbb\xca\x52\xe3\xc1\xd9\xea\x31\xd6\xcb\x15\xea\xe5\x11\x12\xcf\x13\x52\x91\x85\x80\x5a\x76\x5d\x96\xf1\x9e\x6d\x91\x01\x7c\xf9\x2b\xc9\xdf\xb6\x25\xee\x8d\x45\x08\x19\x00\x80\x24\xd7\x28\x71\xd9\x42\xa9\x4c\xf5\xe1\x9b\xde\x56\x19\xed\x51\xfb\x2d\x30\x60\xc9\x72\x14\xf6\x40\x7a\xed\x4d\xb3\x85\xf5\xd7\x5f\x35\xe7\x64\xae\x91\x0e\xb5\xdf\xc2\xd5\xd2\x92\xa3\x92\x14\xf9\xcb\x16\x6a\x92\x12\x75\xb4\x77\x19\xcf\x07\x12\xd9\xb5\xcf\x43\x9a\xd8\x6a\x00\x2e\xa9\x85\x4a\x09\xe7\x76\xbd\x43\x90\x46\xbb\xde\xab\x13\x49\x56\x64\x3d\x36\xaf\x9f\x8b\x9f\xd1\xb6\x54\x21\xbc\x23\x57\x99\x16\xed\x85\xe7\xf5\x73\x91\xdc\x33\x04\x2f\x4a\x85\xeb\x09\x87\xa5\x08\x00\xde\x3b\x16\x51\x30\xc5\x4a\xd4\x0e\xe5\xb0\x2f\x8d\x95\x68\xa7\xad\xf3\x96\x9a\x69\x57\xc7\xd4\x57\xd0\x93\x1a\x97\x00\x21\x58\xa1\x0f\x08\x2b\x7a\x82\xd5\x8b\x29\x61\xbb\x83\xcd\xf7\x5a\xe2\xb9\x2f\x73\xfc\xe1\x8a\x8a\xd9\x16\x80\x8b\x41\x48\x9f\xbd\x98\x72\x1d\x42\x3c\x1b\xbf\xe7\xb8\xe2\xb9\x28\xe0\xf3\x10\x80\x22\x16\xac\x36\xdf\x55\x9e\xda\x98\x27\x0a\x64\x66\x7f\x6f\xbc\x50\xc9\x0c\x22\xc5\x0c\xa2\xf9\x62\x9e\x3f\x9f\x13\x18\xa5\x32\xfa\xc6\x82\x78\x2e\xa9\x1d\xba\x1f\xc2\x6a\xc0\xe9\x4b\x7a\x9f\xd6\xc3\xa1\x4f\x56\xfd\xf7\xbf\x4e\xfd\x76\x0c\x79\x31\xe5\xef\x71\x0e\xd1\x32\x20\xd9\xef\x67\x4d\x99\x11\x4f\xa6\xc9\xc0\xcb\x93\xf7\x46\x0f\x53\x9b\x36\xec\x9a\x37\xf1\xaf\x8c\x52\xa2\x71\x28\xd7\x49\x01\xa5\xd7\xf1\x77\xdd\x58\x3a\x0a\x7b\x61\x85\xab\xcd\x47\x38\x1a\x8b\x3c\x4f\x10\x13\xc3\xbc\x7e\x3b\xae\x43\xf8\x48\xbe\x9e\xfa\xfe\xce\x9a\x26\x2a\x64\x45\x13\x9d\x10\x68\x0f\x07\x0f\x1b\x78\x7e\xf3\x06\xae\xfd\x5d\x48\x7d\xa4\x82\x0a\x8f\x71\x1a\xa0\x1f\x94\x1d\x1b\x07\x53\x1b\x8d\x8b\x8a\x61\x13\x3f\xed\x58\x4b\x2d\x5a\x84\x12\x51\x83\x4c\x04\x9e\x20\xb2\x27\x7d\x00\xa3\xd5\x05\x7c\x8d\xb0\x27\xeb\x7c\xcf\x61\x88\x99\x4e\x0b\x07\x78\x16\xc7\x46\xa1\xdb\x4c\x25\xf6\xdf\xfc\x9e\x30\x6e\x64\xf2\x5f\xce\xd2\x5f\x54\xcd\x7d\x14\xc3\x5c\x1c\x12\xbd\x20\xe5\xd8\x7c\x8a\xb8\xb7\x37\x43\xe5\xeb\x62\xbc\x29\x50\xc2\x0f\xa2\x44\xe5\x78\xee\xeb\x3f\x87\x25\x3d\x3f\x08\xe1\xf9\x1c\x39\x7a\x51\xc8\x39\xbb\xd2\xc8\xcb\x30\x27\xf3\x71\x18\xb4\x31\x76\xfb\x46\xab\x77\x09\x4f\x45\xaa\x9e\x09\x5b\xfa\x07\xcd\x27\x5f\x3f\x64\xd7\xf2\x12\xf5\xcd\x8f\xa2\x81\x45\x92\x01\xf9\xa4\x26\x64\x72\x7e\x4d\x5a\x91\xc6\xa9\xef\xbd\xad\x5f\xaf\xe3\xe8\xdc\xf6\x7f\x9e\x7d\x98\xf4\x9e\xc3\x13\xac\x5a\xa1\x4e\x18\xa9\x0c\xac\xee\xe4\x86\x7b\x17\xde\xcc\xe5\x1a\xa1\x17\x65\x43\xff\xf7\x3a\x8f\x63\xc9\x5d\x17\x1f\xc7\x94\xb2\xeb\x18\xcf\xe3\xc9\xfb\xb0\xcb\x0b\x6e\xce\x1f\x95\xc3\xbb\x24\x1f\x52\x7c\x48\x50\xe2\x5e\x9c\x94\x67\x45\x6c\xd7\x63\x36\x9f\xe0\x32\x9b\xa7\x79\xfc\xfc\x45\x49\x16\x2f\xff\xad\x52\xfe\xd7\xc7\xb2\xa3\xff\x58\x1f\xaf\x23\x39\x69\x64\x78\x20\x5e\x49\xea\x95\x42\x59\xde\x5c\xb7\x38\x3c\x1f\x6e\xae\x69\x1b\xaf\xe4\xc5\xdb\x9f\x8e\x01\x6a\x99\x54\x33\xfe\x4b\x30\x42\xfd\x11\x00\x00\xff\xff\x99\xd2\x3b\x01\x37\x0b\x00\x00"),
		},
		"/status.html": &vfsgen۰CompressedFileInfo{
			name:             "status.html",
			modTime:          time.Date(2018, 6, 25, 7, 35, 55, 128143670, time.UTC),
			uncompressedSize: 2646,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x56\x41\x6f\xe2\x3c\x10\xbd\xf3\x2b\xe6\xf3\xb1\xfa\x20\x52\x8f\x9f\x42\xa4\x96\x7e\x4b\x57\xea\xaa\x15\x94\xed\x5e\x4d\x3c\x21\x56\x83\x8d\xc6\x0e\xdd\xca\xf2\x7f\x5f\x39\x24\x29\x59\x91\xd2\xc2\x61\xf7\x02\x79\xf1\xcb\x7b\x2f\x33\xce\x24\xce\x09\xcc\xa4\x42\x60\x39\x72\xc1\xbc\x8f\xff\x19\x0e\x41\xc9\x9f\x30\x1c\x26\xce\xa1\x12\xde\x0f\x06\x6f\xac\x54\x2b\x8b\xca\x32\xef\x07\x00\xb1\x90\x5b\x48\x0b\x6e\xcc\xb8\x5a\xe0\x52\x21\x0d\xb3\xa2\x94\x82\x25\x03\x00\x80\x38\xbf\x04\x29\xc6\x8c\x4a\x65\xe5\x1a\x59\x32\xdb\x1d\xc0\x57\x95\x69\x5a\x73\x2b\xb5\x8a\xa3\xfc\xb2\x66\x5b\xbe\x2c\xb0\x51\xdc\x81\xea\x77\x98\x6a\x25\x50\x19\x14\x35\x5e\x6a\x12\x48\x2d\x34\x96\xe4\xa6\x45\xb9\xde\x22\xd5\x01\x82\xe8\x52\x8b\xd7\x06\x05\x4c\x6f\x20\xc0\x3c\x59\x6c\x42\xa6\x38\xb2\x79\x77\x45\x24\xce\x8d\xae\x25\xd9\x7c\xb4\x78\x9c\x78\x1f\x47\x56\xec\x09\x45\xfb\x4a\x07\x64\x9f\x34\x3d\x4b\xb5\x82\x1b\x49\x98\x5a\x4d\xaf\x3d\x0e\x93\xa7\x9b\xa3\xda\xce\xc9\x0c\x94\xb6\x30\x9a\x61\xa1\xb9\x98\x68\x95\xc9\xd5\xbc\x4c\x53\x34\xc6\xfb\xa6\x66\x82\xab\x15\x12\xab\x1b\xf7\x7b\xa0\xdd\x45\x25\x55\x65\x07\xaa\x84\x0e\x67\x92\x59\x8f\x51\x7d\x90\x95\x85\x73\x58\x18\xf4\xfe\x0b\x97\x05\x8a\xda\xf1\x93\x15\xba\xe3\xc6\x82\x69\x25\x21\xfd\x70\xc0\x51\xb8\x74\x97\xed\x51\xae\xf1\xa4\xfe\xdc\x22\x17\x90\xe6\xa5\x7a\x36\x7d\x9d\x09\x8b\x13\x5d\x2a\x7b\x9a\x78\xb5\xd5\x0d\x92\xc4\x3e\x87\x10\x7e\x5e\x11\x4e\xb3\x79\xba\xba\x83\x54\x13\x95\x9b\x50\xb2\xde\xfb\x68\x19\xa7\xb9\x4c\x35\xe9\xd2\x4a\xd5\x7b\x1b\x2d\xe1\x44\xfd\xfb\x6f\x57\x3f\x1e\x66\xf7\x93\x79\x9f\x7e\x4b\x38\x41\x7b\x3a\xe9\x55\x9d\xbe\xbb\x6b\xe2\x68\x6f\x76\xc4\x51\x35\x5d\x92\x41\x67\xb0\x2d\x4b\x59\x08\xf9\x36\xcc\x58\x72\x1d\xce\xfc\x55\xf3\x0d\x4c\xaa\x37\x38\x66\xa4\x5f\x58\xf2\x1d\xc9\x54\xa1\x0e\x56\xa4\x5e\x6d\xfe\x3f\x5b\xec\x8e\xd3\x0c\xb7\xf2\x03\x56\x0d\xed\x2c\xaf\x6b\xe2\x2a\xcd\x8f\x38\xed\x48\xe7\xf9\x84\xe6\x2e\x0c\xd2\x31\xab\x86\x77\xbe\xdb\x0d\xb7\x7d\xaf\xa6\x8e\x5b\xe0\x9d\xe5\x36\xd5\x1f\xdb\x1b\x2d\xef\xcc\x47\x87\x17\x48\x76\xcd\x15\x5f\x21\x19\x96\x5c\xed\xc3\x3f\xfb\xcc\x54\x73\xe3\x7f\x25\x36\x5a\x2a\xdb\xad\x46\xb7\xa2\xce\x51\x78\xe1\xc2\xa8\x13\xbe\xfa\x36\x3a\x28\xec\x5c\x74\x01\xfb\x5c\x58\xcc\xee\x0c\xf0\xe2\x85\xbf\x1a\xc8\xf9\x16\x61\x9e\xe6\xb8\xc6\x7f\xe1\x56\x1b\x0b\x5c\x09\x78\xe0\xa1\x4f\x68\xe1\x22\xda\x13\x6e\xbb\xb2\xe3\x7b\xff\x5f\x14\xc5\x1c\x72\xc2\x6c\xcc\xba\xa7\x9d\x1b\x05\x31\xef\x59\xd2\x1e\xc6\x11\x0f\x20\x68\xbf\xbf\x67\x9a\x4f\xc0\x23\x5d\x8d\x23\x21\xb7\xc9\xa0\x61\xff\x0a\x00\x00\xff\xff\xe5\xe4\xe6\xaa\x56\x0a\x00\x00"),
		},
		"/targets.html": &vfsgen۰CompressedFileInfo{
			name:             "targets.html",
			modTime:          time.Date(2018, 6, 25, 7, 35, 55, 128143670, time.UTC),
			uncompressedSize: 3086,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x56\x5f\x8f\xdb\x36\x0c\x7f\xcf\xa7\x20\xbc\x60\xd8\x80\x3a\x46\xfb\xd8\xd9\x1e\xf6\xa7\x40\x07\x74\x43\xd7\xeb\xed\x61\x2f\x85\x6c\x31\xb1\xee\x14\xc9\x93\xe8\xec\x02\x57\xdf\x7d\x90\xe4\x7f\x49\x2e\xb7\xde\x1e\xf6\x12\x84\x12\x49\xfd\x48\xfe\x48\xba\xef\x39\x6e\x85\x42\x48\x1a\x64\x3c\x71\x6e\x95\x4b\xa1\xee\x81\x8e\x2d\x16\x09\xe1\x03\x65\xb5\xb5\x09\x18\x94\x45\x62\xe9\x28\xd1\x36\x88\x94\x40\x63\x70\x5b\x24\x7d\x0f\x2d\xa3\xe6\xbd\xc1\xad\x78\x00\xe7\x32\x4b\x8c\x44\xed\x6d\x32\x62\x66\x87\x64\x37\xb5\xb5\xdf\x1f\x8a\xbe\x87\xaa\x13\x92\xff\x81\xc6\x0a\xad\xc0\xb9\xa4\x5c\xe5\xb6\x36\xa2\x25\xb0\xa6\xbe\xee\xeb\x6e\x76\x75\x77\xcd\x53\x9e\x45\x4f\xe5\xaa\xef\x51\x71\xe7\x56\xab\xd5\x1c\x5a\xad\x15\xa1\x22\x1f\x1d\x40\xce\xc5\x01\x6a\xc9\xac\x2d\xc2\x05\x13\x0a\x4d\xba\x95\x9d\xe0\x49\xb9\x02\x00\xc8\x9b\x97\xe5\xc7\xf8\x62\x9e\x35\x2f\x87\x43\x6f\x26\x78\x91\xd8\x46\xff\x3d\xdc\x26\xa3\x9f\x8a\x54\xba\x33\xba\x6b\x13\xe0\x8c\x58\x4a\x7a\xb7\x93\x58\x24\x55\x47\xa4\x95\x1d\xfc\x02\xe4\x92\x55\x28\x17\x56\xe0\x2d\x5b\x23\xf6\xcc\x1c\x81\xd5\x24\x0e\x38\x29\x03\xe4\x42\xb5\x1d\x0d\xb5\x30\x8c\x0b\x9d\x80\x62\x7b\x5f\x98\x11\x80\x47\xc4\xa4\x4c\xa7\x03\xd6\x91\xae\xf5\xbe\x95\x48\x58\x24\x7a\xbb\x4d\xa0\x6e\xb0\xbe\x47\x5e\xc2\x0f\x52\x8e\x40\xb2\x80\xe4\x8b\x70\x3d\x1f\x50\xa7\x1a\x64\x92\x9a\xe3\x53\xb0\x4a\xb8\x1d\xd5\x2e\x41\xe5\x19\x17\x87\x72\x15\x2e\xfa\xde\x30\xb5\x43\x58\xdf\xe9\xea\x05\xac\x5b\xad\x25\xbc\x2e\x60\x13\x8b\xf0\x5e\x6b\x69\x43\x61\xbd\xe6\x7a\xf0\xe8\x15\x54\xb7\x7f\x3b\x48\xc1\x68\x56\x22\x4d\x2c\xf8\x90\xa8\xa6\xbb\xb9\xca\x43\x1a\x88\x55\x12\xd3\x89\x22\x73\x11\x9b\x57\xa3\xca\x9d\xae\x3e\xf9\xbe\x41\xd3\xf7\x62\x0b\x92\x60\x02\x10\x1f\x71\x0e\xb8\x07\x6f\x06\x5a\x2e\x73\xc9\x42\xae\xee\x74\x95\xf6\xbd\x8f\xcd\xb9\xb1\xab\xbe\x3a\x39\x2c\xc7\x7f\xf0\xcd\x1c\xa0\x73\xd9\x18\x88\x73\xd0\xb5\xdf\xe6\x19\x5b\xf8\x8e\xc4\x1b\x0a\x15\x85\x64\x8e\x2b\xd4\x04\xf0\xa1\x65\x8a\x23\x4f\x43\xa0\x70\x51\x75\x4f\x74\x90\x68\x6d\x9e\x45\x0f\x53\x02\xb2\xe6\xd5\xf4\x3f\x1a\x2f\x53\x06\x53\xe2\x38\x2a\x8b\x7c\x90\x2b\x6d\x38\x9a\x49\xb4\x64\x44\x3b\x49\x8d\x3e\x2c\x32\xec\xdd\xfa\xb4\x2e\xd3\xcc\x91\x98\x90\x76\xa1\xe3\xb5\xcc\x52\x0c\x66\xe5\x1b\xc5\x5b\x2d\x14\xe5\x19\x35\x97\xb7\x37\xc4\x08\x1f\xbf\x7a\xe7\xc9\x67\xaf\xdd\x59\x82\x9b\xda\xb0\xf6\x8a\xf1\x1b\x63\xb4\x39\xbf\xca\xb3\x25\x40\x7f\x8b\x8c\x2f\x83\xac\x34\x3f\xce\xf2\x44\xf4\x05\x5b\xaf\x06\x3a\x25\x07\x87\x78\x93\x53\x8d\x40\xb0\x71\x48\x6f\x6e\x3f\xbc\x83\xcf\xb0\x93\xba\x62\xf2\xf6\xc3\xbb\xc8\x2a\x7f\xba\xb9\xa9\x1b\xdc\xa3\x73\xaf\xb3\x6c\x38\x79\xab\x2d\x39\x37\x08\xef\x19\x35\xce\x79\x72\xe5\x95\x39\x7f\x61\x02\x1c\xfa\xf6\x05\xac\x0f\x4c\x76\x68\x43\x77\x7a\xe3\xdf\x3b\x34\x47\x38\x89\xe4\xcc\x50\x8c\x46\xde\x66\x30\x7f\x44\x1f\x20\xb7\x2d\x53\x63\xc8\x71\x5a\x85\xdf\x99\xae\x7d\x1f\x61\x38\xe7\x23\x8e\xbe\x9c\x4b\xf2\xcc\x5b\x9e\x23\xf7\x10\xe2\x9e\xf8\xf7\xd3\x3c\x23\x7e\x35\xfb\x7e\x47\xe1\x65\xea\x97\x68\x99\x44\x43\x10\x7e\xd3\xbe\x87\x4d\x1c\x4a\xf0\x19\x62\x2b\x7f\xd4\x3f\x79\x3d\x70\x0e\x82\xb3\x4f\x42\x71\x51\x33\xd2\x06\xfc\xfa\x4d\xbb\xb6\x45\x53\x33\x7b\xf9\x8a\x87\x3b\x78\xbb\x08\xe4\xb1\xb0\x9f\x0c\x24\xe4\xce\x3e\x1d\x49\xdd\x19\xab\x4d\x1a\xe8\x86\xe6\x6c\xd7\x91\xd6\x92\x44\x9b\x00\x09\xf2\xf2\x70\xdd\xd0\x5e\x16\x64\x3a\x8c\xa2\x36\x62\x27\x14\x93\xe9\xa0\x95\x57\xe5\x8f\xb8\xd5\x06\xfd\x17\x86\x87\x20\xd4\xee\x75\x9e\x55\xe5\x44\x91\x7b\x4f\x91\x40\xa9\x9f\x85\xad\xfd\x9c\x40\x1e\x5b\x75\xf3\x2b\x6b\x9d\xf3\xbc\xec\xfb\x35\x1e\x22\x85\x7c\x5e\x69\x2f\x3d\x81\xd7\xf7\xce\x15\x5f\xff\xd5\x69\xfa\x2e\x28\x38\x37\x0a\xe7\xb3\x78\xce\x67\xe4\x50\xa0\x70\x98\x50\xf1\x25\xd8\xcc\x2f\x82\x1f\x47\xc9\x53\xa4\x3e\xe9\x86\x80\x2a\x3a\xfd\xdf\x88\x2d\x2d\x3e\xef\x2d\x8e\x5b\xd6\x49\x4a\x4a\xa5\x15\x3e\xaf\x63\xfe\x03\xd1\x2c\xa5\x36\x8c\xd2\x8b\x0a\x84\x2d\xba\xf1\xc3\x36\xce\xda\xcd\x2f\xf6\x4f\x34\xda\xb9\xdf\xf0\x10\x76\x68\x88\xac\xef\xad\x50\x35\x2e\x15\x9d\x03\xb6\xd3\xcf\xee\x5e\xf4\x53\xfb\x92\xf4\x33\x8c\x30\xd6\x2f\x63\xbe\xd6\xde\x71\xd7\x9f\xf7\x71\x98\xb5\x0b\x6f\x8f\x27\xf8\x4b\xc0\x9f\x2e\x94\x73\x8b\x3c\x3b\x59\x28\x79\x16\x76\xeb\xf0\xf9\x1a\x3f\xa8\xa2\x19\xa0\xe2\x71\x2c\x0f\xc7\xa3\xa7\x7f\x02\x00\x00\xff\xff\x5b\x88\x7f\x6f\x0e\x0c\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/_base.html"].(os.FileInfo),
		fs["/alerts.html"].(os.FileInfo),
		fs["/config.html"].(os.FileInfo),
		fs["/flags.html"].(os.FileInfo),
		fs["/graph.html"].(os.FileInfo),
		fs["/rules.html"].(os.FileInfo),
		fs["/service-discovery.html"].(os.FileInfo),
		fs["/status.html"].(os.FileInfo),
		fs["/targets.html"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr: gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
