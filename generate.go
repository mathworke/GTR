package main

import (
	"GTR/assets"
	"GTR/ui"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
)

type Data struct {
	Module        string              `json:"module"`
	ModuleVersion string              `json:"module_version"`
	DateTesting   string              `json:"date_testing"`
	DateVersion   string              `json:"date_version"`
	Post          string              `json:"post"`
	Tester        string              `json:"tester"`
	UsingProject  string              `json:"using_project"`
	MainTask      string              `json:"main_task"`
	TaskChanged   []string            `json:"taskChanged"`
	TestCases     []map[string]string `json:"testCases"`
	Bugs          []map[string]string `json:"createBug"`
	Comment       string              `json:"comment"`
	Release       string              `json:"release"`
}

type Report struct {
	Modules  *ui.ModuleInformation
	Task     *ui.Task
	TestCase *ui.SelectTestCase
	Bugs     *ui.CreateBug
}

func (r *Report) Generate(comment string, release bool, person *assets.Person, logger *assets.Logger) {
	logger.LogIngo("starting generate data.json")

	// Setup
	r.Task.FillChanged(logger)
	r.Bugs.FillBugs(logger)

	data := Data{
		Module:        r.Modules.Module,
		ModuleVersion: r.Modules.ModuleVersion,
		DateTesting:   r.Modules.DateTesting,
		DateVersion:   r.Modules.DateFirmware,
		Post:          person.Post,
		Tester:        person.Fio,
		UsingProject:  r.Modules.UseProject,
		MainTask:      r.Task.MainTask,
		TaskChanged:   *r.getTaskChanged(),
		TestCases:     *r.getTestCase(),
		Bugs:          *r.getBugs(),
		Comment:       comment,
	}

	if release {
		data.Release = "рекомендую"
	} else {
		data.Release = "не рекомендую"
	}

	// create config file for report
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		logger.PANIC(err.Error())
	}

	err = os.WriteFile("assets/data.json", file, 0644)
	if err != nil {
		logger.PANIC(err.Error())
	}

	// run transforms report
	cmd := exec.Command("python", "assets/generate_reports.py")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		logger.PANIC(err.Error())
	}

	filename := make(chan string, 1)

	receiveQuestion := func(stdout io.ReadCloser) {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			filename <- scanner.Text()
		}
	}

	go receiveQuestion(stdout)

	cmd.Start()
	cmd.Wait()

	logger.LogIngo("report generated")

	msg := <-filename
	if string(msg[:3]) == "err" {
		logger.PANIC("Error completing python script\n%v\n", msg)
	} else {
		cmd = exec.Command("cmd", fmt.Sprintf("/C start docs/%v", msg))
		cmd.Run()
	}
}

func (r *Report) getTaskChanged() *[]string {
	return &r.Task.Changed
}

func (r *Report) getTestCase() *[]map[string]string {
	testCases := make([]map[string]string, 0)
	for _, pass := range r.TestCase.CompleteTestCase {
		for _, item := range pass {
			delta := len(testCases)
			testCases = append(testCases, make(map[string]string))

			testCases[delta]["label"] = item.Номер
			if item.Пройден {
				testCases[delta]["status"] = "Пройден"
			} else {
				testCases[delta]["status"] = "Не пройден"
			}
			testCases[delta]["tester"] = item.Ответственный
			testCases[delta]["comment"] = item.Комментарий

		}
	} // [0] - passed; [1] - failed

	/*
		for _, item := range r.TestCase.CompleteTestCase[1] {
			delta := len(testCases)
			testCases = append(testCases, make(map[string]string))

			testCases[delta]["label"] = item.Номер
			if item.Пройден {
				testCases[delta]["status"] = "Пройден"
			} else {
				testCases[delta]["status"] = "Не пройден"
			}
			testCases[delta]["tester"] = item.Ответственный
			testCases[delta]["comment"] = item.Комментарий

		} // test failes
	*/
	return &testCases
}

func (r *Report) getBugs() *[]map[string]string {
	bugs := make([]map[string]string, 0)

	for _, item := range r.Bugs.Bugs {
		delta := len(bugs)
		bugs = append(bugs, make(map[string]string))

		bugs[delta]["number"] = item.Номер
		bugs[delta]["name"] = item.Название
		if item.Исправленность {
			bugs[delta]["done"] = "Исправлено"
		} else {
			bugs[delta]["done"] = "Не исправлено"
		}
		bugs[delta]["critical"] = item.Серьёзность
	}

	return &bugs
}
