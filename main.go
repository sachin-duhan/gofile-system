package main

import (
    "fmt"
    "io"
    "os"
    "path/filepath"
    "time"
)

// File represents a file in the file system.
type File struct {
    Name     string
    Path     string
    Contents string
    Version  int
}

// FileSystem represents the file system.
type FileSystem struct {
    files map[string]*File
}

// NewFileSystem initializes a new file system.
func NewFileSystem() *FileSystem {
    return &FileSystem{
        files: make(map[string]*File),
    }
}

// CreateFile creates a new file in the file system.
func (fs *FileSystem) CreateFile(name, path, contents string) error {
    _, exists := fs.files[filepath.Join(path, name)]
    if exists {
        return fmt.Errorf("file %s already exists", name)
    }

    fs.files[filepath.Join(path, name)] = &File{
        Name:     name,
        Path:     path,
        Contents: contents,
        Version:  1,
    }

    return nil
}

// DeleteFile deletes a file from the file system.
func (fs *FileSystem) DeleteFile(name, path string) error {
    key := filepath.Join(path, name)
    _, exists := fs.files[key]
    if !exists {
        return fmt.Errorf("file %s does not exist", name)
    }

    delete(fs.files, key)
    return nil
}

// Copy copies a file to a given path in the file system.
func (fs *FileSystem) Copy(name, srcPath, destPath string) error {
    srcKey := filepath.Join(srcPath, name)
    file, exists := fs.files[srcKey]
    if !exists {
        return fmt.Errorf("file %s does not exist", name)
    }

    destKey := filepath.Join(destPath, name)
    _, exists = fs.files[destKey]
    if exists {
        return fmt.Errorf("file %s already exists in destination path", name)
    }

    copiedFile := *file
    copiedFile.Path = destPath
    fs.files[destKey] = &copiedFile
    return nil
}

// MoveFile moves a file to a given path in the file system.
func (fs *FileSystem) MoveFile(name, srcPath, destPath string) error {
    srcKey := filepath.Join(srcPath, name)
    file, exists := fs.files[srcKey]
    if !exists {
        return fmt.Errorf("file %s does not exist", name)
    }

    destKey := filepath.Join(destPath, name)
    _, exists = fs.files[destKey]
    if exists {
        return fmt.Errorf("file %s already exists in destination path", name)
    }

    file.Path = destPath
    delete(fs.files, srcKey)
    fs.files[destKey] = file
    return nil
}

// SaveVersion saves a new version of the file in the file system.
func (fs *FileSystem) SaveVersion(name, path, contents string) error {
    key := filepath.Join(path, name)
    file, exists := fs.files[key]
    if !exists {
        return fmt.Errorf("file %s does not exist", name)
    }

    file.Version++
    file.Contents = contents
    return nil
}

// SwitchVersion switches to an older version of the file.
func (fs *FileSystem) SwitchVersion(name, path string, version int) error {
    key := filepath.Join(path, name)
    file, exists := fs.files[key]
    if !exists {
        return fmt.Errorf("file %s does not exist", name)
    }

    if version <= 0 || version > file.Version {
        return fmt.Errorf("invalid version number")
    }

    // Simulate version switching by printing the old version
    fmt.Printf("Old version of %s (Version %d):\n%s\n", name, version, file.Contents)
    return nil
}

func main() {
    fs := NewFileSystem()

    // Create a file
    err := fs.CreateFile("example.txt", "/documents", "This is a sample file.")
    if err != nil {
        fmt.Println(err)
        return
    }

    // Save a new version of the file
    err = fs.SaveVersion("example.txt", "/documents", "This is an updated version of the file.")
    if err != nil {
        fmt.Println(err)
        return
    }

    // Switch to an older version of the file
    err = fs.SwitchVersion("example.txt", "/documents", 1)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Copy the file to a different path
    err = fs.Copy("example.txt", "/documents", "/backup")
    if err != nil {
        fmt.Println(err)
        return
    }

    // Move the file to a different path
    err = fs.MoveFile("example.txt", "/documents", "/archive")
    if err != nil {
        fmt.Println(err)
        return
    }

    // Delete the file
    err = fs.DeleteFile("example.txt", "/archive")
    if err != nil {
        fmt.Println(err)
        return
    }
}
