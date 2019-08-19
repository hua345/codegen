package models

import (
	"codegen/pkg/fileUtil"
	"codegen/pkg/templateUtil"
	"path"
)

type FileMapDto struct {
	TplSrcPath string `json:"tplSrcPath"`
	TplDstPath string `json:"tplDstPath"`
}

/**
 * 添加子目录文件模板映射关系
 */
func appendSubDirTemplateList(templateConfigPath, codeConfigPath string, fileMapDtoList []FileMapDto) []FileMapDto {
	subDirList := templateUtil.GetTemplateSubDirList(templateConfigPath)
	// subDirList := util.GetSubDirList(templateConfigPath)
	for _, value := range subDirList {
		templateSubPath := path.Join(templateConfigPath, value)
		codeSubDir := path.Join(codeConfigPath, value)
		fileUtil.CheckDirAndMkdir(codeSubDir)
		fileMapDtoList = appendTemplateList(templateSubPath, codeSubDir, fileMapDtoList)
	}
	return fileMapDtoList
}

/**
 * 添加Mybatis文件模板映射关系
 */
func appendMybatisTemplateList(mybatisTemplatePath, mybatisPath string, fileMapDtoList []FileMapDto) []FileMapDto {
	// go-bindata
	mybatisFileNameList := templateUtil.GetTemplateFilesName(mybatisTemplatePath)
	// mybatisFileNameList := util.GetFilesName(mybatisTemplatePath)
	for _, value := range mybatisFileNameList {
		if fileUtil.GetFileSuffix(value) == ".jar" {
			// 如果是jar文件直接拷贝
			templateUtil.CopyTemplateFile(path.Join(mybatisPath, value),
				path.Join(mybatisTemplatePath, value))
		} else {
			fileMapDtoList = append(fileMapDtoList,
				FileMapDto{path.Join(mybatisTemplatePath, value),
					path.Join(mybatisPath, value)})
		}
	}
	return fileMapDtoList
}

/**
 * 添加文件模板映射关系
 */
func appendTemplateList(templatePath, codePath string, fileMapDtoList []FileMapDto) []FileMapDto {
	// go-bindata
	templateFileNameList := templateUtil.GetTemplateFilesName(templatePath)
	// templateFileNameList := util.GetFilesName(templatePath)
	for _, value := range templateFileNameList {
		fileMapDtoList = append(fileMapDtoList,
			FileMapDto{path.Join(templatePath, value),
				path.Join(codePath, value)})
	}
	return fileMapDtoList
}
