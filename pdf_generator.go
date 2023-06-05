package go_latex

import (
	"io/ioutil"
	"os"
	"os/exec"
)

// GeneratePDF generates a PDF document from LaTeX source
func GeneratePDF(latexSource string) error {
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
