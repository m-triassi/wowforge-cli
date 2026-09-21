package curseforge

import (
	"archive/zip"
	"os"
	"path/filepath"
	"testing"
)

type testFile struct {
	name    string
	content string
}

func buildTestZip(t *testing.T, destPath string, testFiles []testFile, withDirEntries bool) {
	t.Helper()

	zf, err := os.Create(destPath)
	if err != nil {
		t.Fatalf("failed to create zip file: %v", err)
	}
	defer zf.Close()

	zw := zip.NewWriter(zf)

	if withDirEntries {
		seen := map[string]bool{}
		for _, file := range testFiles {
			dir := filepath.Dir(file.name)
			for dir != "." && dir != "/" && !seen[dir] {
				seen[dir] = true
				hdr := &zip.FileHeader{Name: dir + "/"}
				hdr.SetMode(os.ModeDir | 0o755)
				if _, err := zw.CreateHeader(hdr); err != nil {
					t.Fatalf("failed to write dir entry %q: %v", dir, err)
				}
				dir = filepath.Dir(dir)
			}
		}
	}

	for _, file := range testFiles {
		w, err := zw.Create(file.name)
		if err != nil {
			t.Fatalf("failed to create zip entry %q: %v", file.name, err)
		}
		if _, err := w.Write([]byte(file.content)); err != nil {
			t.Fatalf("failed to write zip entry %q: %v", file.name, err)
		}
	}

	if err := zw.Close(); err != nil {
		t.Fatalf("failed to close zip writer: %v", err)
	}
}

func TestClient_InstallAddon(t *testing.T) {
	testFiles := []testFile{
		{name: "MyAddon/MyAddon.toc", content: "toc-content"},
		{name: "MyAddon/Libs/Lib.lua", content: "lib-content"},
	}

	tests := []struct {
		name           string
		withDirEntries bool
	}{
		{name: "flat zip without directory entries", withDirEntries: false},
		{name: "zip with directory entries", withDirEntries: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpDir := t.TempDir()
			zipPath := filepath.Join(tmpDir, "addon.zip")
			buildTestZip(t, zipPath, testFiles, tt.withDirEntries)

			destDir := filepath.Join(tmpDir, "dest")
			if err := os.MkdirAll(destDir, 0o755); err != nil {
				t.Fatalf("failed to create dest dir: %v", err)
			}

			client := NewClient()
			file := File{Location: zipPath}

			if err := client.InstallAddon(file, destDir); err != nil {
				t.Fatalf("InstallAddon returned an error: %v", err)
			}

			for _, file := range testFiles {
				installedPath := filepath.Join(destDir, file.name)
				got, err := os.ReadFile(installedPath)
				if err != nil {
					t.Fatalf("expected file %q to exist: %v", installedPath, err)
				}
				if string(got) != file.content {
					t.Errorf("file %q content = %q, want %q", installedPath, got, file.content)
				}
			}
		})
	}
}
