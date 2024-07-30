# GTR - генератора отчёт по тестированию модуделй ввода-вывода серии MKLogic500

### Зависимости
1. go version go1.22.4
2. gcc-tdm-x86_64
3. python 3.12.4 *( pip install -r requerements.txt )*
---
### Конфигурация файлов
**Папка с файлами должна находится в одно, и том же месте где исполняемый файл. Обязательное название конфигурационной папки - *"config/"***
1. Конфигурационный файл ***module.xml*** - содержит информацию о модулях серии ***MKLogic500***. Пример заполнения файла:
```xml
<modules> <!-- Список модулей -->
    <module type="MK-501-022"> <!-- Название модуля -->
        <description>Модуль центрального процессора</description> <!-- Описание модуля -->
            <tests> <!-- Список тест-кейсов для данного модуля -->
                <test_case>
                    <id>1-123</id> <!-- ID тест-кейса -->
                    <title>Тестовый тест-кейс #1</title> <!-- Название тест-кейса -->
                </test_case>
                <test_case>
                    <id>1-456</id>
                    <title>Тестовый тест-кейс #2</title>
                </test_case>
                <test_case>
                    <id>1-789</id>
                    <title>Тестовый тест-кейс #3</title>
                </test_case>
            </tests>
    </module>
</modules>
```
---
2. Конфигурационный файл ***person*** - содержит информация о тестировщике, проводящий тестирование для выбранного модуля. Пример заполнения файла:
```xml
<person>
    <fio>Иванов И.И.</fio> <!-- Инициалы тестировщика проводяшего тестирование -->
    <post>Инженер-программсит 2-ой категории</post> <!-- Должность тестироващика -->
</person>
```
---
3. Конфигурационный файл ***appearance-settings.toml*** - содержит базовый настройки приложения. Данный файл, создается автоматически и хранится по пути **%APPDATA%/Cogent Core/**. Пример конфигурации:

```toml
Theme = 'Light' 

[Color]
  R = 0
  G = 153
  B = 130
  A = 255

[Screens]
  [Screens.20]
    Zoom = 0.0

  [Screens.'Generic PnP Monitor']
    Zoom = 100.0 

```
*Примечание: если нужны более гибкие настйроки внешнего вида приложения - нужнен специальный исполняемый файл. Могу собрать исполняемый файл и скинуть*
---