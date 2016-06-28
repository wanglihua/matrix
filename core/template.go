package core

import (
    htmlTemplate "html/template"
    textTemplate "text/template"
    "bytes"
)

func FormatText(templateName string, templateContent string, args  map[string]interface{}) string {
    template, err := textTemplate.New(templateName).Parse(templateContent)
    if err != nil {
        panic(err)
    }

    var buffer bytes.Buffer
    template.Execute(&buffer, args)
    return buffer.String()
}

func FormatHtml(templateName string, templateContent string, args map[string]interface{}) htmlTemplate.HTML {
    template, err := textTemplate.New(templateName).Parse(templateContent)
    if err != nil {
        panic(err)
    }

    var buffer bytes.Buffer
    template.Execute(&buffer, args)
    return htmlTemplate.HTML(buffer.String())
}
