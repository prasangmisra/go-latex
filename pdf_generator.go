package go_latex

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
)

// GeneratePDF generates a PDF document from LaTeX source
func GeneratePDF(latexSource string) error {
	// Check if LaTeX is installed
	if !isLaTeXInstalled() {
		// If not installed, try starting the Docker container
		if err := startLaTeXDockerContainer(); err != nil {
			return err
		}
	}

	// ... Rest of the function code ...
	// Create a temporary file with .tex extension
	tmpfile, err := ioutil.TempFile("", "*.tex")
	if err != nil {
		return err
	}
	defer os.Remove(tmpfile.Name())

	// Write the LaTeX source to the temporary file
	if _, err := tmpfile.WriteString(latexSource); err != nil {
		return err
	}

	// Close the temporary file
	if err := tmpfile.Close(); err != nil {
		return err
	}

	// Execute LaTeX command to generate PDF
	cmd := exec.Command("pdflatex", "-interaction=batchmode", tmpfile.Name())
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

// isLaTeXInstalled checks if LaTeX is installed on the system
func isLaTeXInstalled() bool {
	cmd := exec.Command("pdflatex", "--version")
	err := cmd.Run()
	return err == nil
}

// startLaTeXDockerContainer starts a Docker container with a LaTeX image
func startLaTeXDockerContainer() error {
	// Check if Docker is installed
	if !isDockerInstalled() {
		return errors.New("docker is not installed")
	}

	// Pull the LaTeX Docker image
	cmd := exec.Command("docker", "pull", "texlive/texlive")
	if err := cmd.Run(); err != nil {
		return err
	}

	// Run the LaTeX Docker container in the background
	cmd = exec.Command("docker", "run", "-d", "texlive/texlive")
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

// isDockerInstalled checks if Docker is installed on the system
func isDockerInstalled() bool {
	cmd := exec.Command("docker", "--version")
	err := cmd.Run()
	return err == nil
}
