package services

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/wberdnik/CICD_HTTP_Agent/internal/config"
)

func HandlerAgent(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		Http404Exception(w, r)
		return
	}

	query := r.URL.Query()
	project := query.Get("project")

	if len(project) == 0 || strings.Contains(project, "/") {
		Http400Exception(w, errors.New("Error project param"))
		return
	}

	scriptName := "/etc/cicd_agent/" + project + ".sh"

	file, err := os.OpenFile(scriptName, os.O_RDONLY, 0644)
	file.Close()

	if err != nil {
		Http404Exception(w, r)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, config.MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(config.MAX_UPLOAD_SIZE); err != nil {
		Http400Exception(w, errors.New("File too big (1Gb)"))
		return
	}

	fileH, fileHeader, err := r.FormFile("file")
	if err != nil {
		Http400Exception(w, err)
		return
	}

	defer fileH.Close()

	err = os.MkdirAll("/tmp/cicd_agent", os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fileName := fmt.Sprintf("/tmp/cicd_agent/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename))

	dst, err := os.Create(fileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(dst, fileH)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dst.Close()

	cmd, err := exec.Command(scriptName, fileName).Output()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "text/plan")
	w.Write([]byte("Success: " + string(cmd)))
}
