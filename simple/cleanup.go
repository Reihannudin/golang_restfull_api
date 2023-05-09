package simple

import "fmt"

//File

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
	fmt.Print("Close File", f.Name)
}

// Connection
type Connection struct {
	File *File
}

func (c *Connection) Close() {
	fmt.Print("Close Connection", c.File.Name)
}

func NewConnection(file *File) (*Connection, func()) {
	connection := &Connection{
		File: file,
	}
	return connection, func() {
		connection.Close()
	}
}
