package _12cleanup_function

import "fmt"

type File struct {
	Name string
}

func NewFile(name string) (*File, func()) {
	file := &File{Name: name}
	return file, func() {
		file.Close()
	}
}

func (f *File) Close() {
	fmt.Println("Close file ", f.Name)
}

type Connection struct {
	*File
}

func NewConnection(file *File) (*Connection, func()) {
	connection := &Connection{File: file}
	return connection, func() {
		connection.Close()
	}
}

func (c *Connection) Close() {
	fmt.Println("Close connection ", c.File.Name)
}
