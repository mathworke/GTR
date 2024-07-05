# генерация отчета по шаблону template.docx. используется jinx2

from docxtpl import DocxTemplate
from json import loads

class GenerateReport():
	def __init__(self, path_template: str, path_param: str):
		self.rep = DocxTemplate('template.docx')

		with open(path_param, 'r', encoding='utf-8') as file:
			file_content = file.read()
			self.data = loads(file_content)

	def CreateReport(self):
		self.rep.render(self.data)
		extencion = '.docx'
		filename = f'Отчет_{self.data["module"]}_{self.data["module_version"]}{extencion}'
		self.rep.save(filename)

if __name__ == "__main__":
	reports = GenerateReport("template.docx", "data.json")
	reports.CreateReport()

