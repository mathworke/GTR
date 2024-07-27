#!/usr/local/bin/python

from docxtpl import DocxTemplate
from json import loads

class GenerateReport():
	def __init__(self, path_template: str, path_param: str):
		self.rep = DocxTemplate(path_template)

		with open(path_param, 'r', encoding='utf-8') as file:
			file_content = file.read()
			self.data = loads(file_content)

	def CreateReport(self) -> str:
		self.rep.render(self.data)
		extencion = '.docx'
		filename = f'Report_{self.data["module"]}_{self.data["module_version"]}{extencion}'
		self.rep.save("docs/"+filename)
		return filename

if __name__ == "__main__":
	reports = GenerateReport("template.docx", "data.json")
	print(reports.CreateReport())

