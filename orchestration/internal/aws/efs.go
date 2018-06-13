package aws

import (
	"github.com/aws/aws-sdk-go/service/efs"
	"strings"
)

func GetEFSId(projectId string) string {
	fsList := getAllEFS()

	tags := make(map[string]string, 2)
	tags["ProjectId"] = projectId
	tags["Type"] = "persistence"

	fsList = filterByTags(fsList, tags)

	if len(fsList) < 1 {
		panic("No EFS FileSystem found.")
	}

	return *(fsList[0].FileSystemId)
}

func getAllEFS() []*efs.FileSystemDescription {
	input := &efs.DescribeFileSystemsInput{}

	result, err := EFSClient.DescribeFileSystems(input)
	if err != nil {
		panic(err)
	}

	return result.FileSystems
}

func filterByTags(fsList []*efs.FileSystemDescription, tags map[string]string) []*efs.FileSystemDescription {
	result := make([]*efs.FileSystemDescription, 0)

	for _, fs := range fsList {
		fsTags := getTags(fs)

		if matchMaps(tags, fsTags) {
			result = append(result, fs)
		}
	}

	return result
}

func getTags(fs *efs.FileSystemDescription) map[string]string {
	input := &efs.DescribeTagsInput{FileSystemId:fs.FileSystemId}

	result, err := EFSClient.DescribeTags(input)
	if err != nil {
		panic(err)
	}

	tags := make(map[string]string, len(result.Tags))

	for _, tag := range result.Tags {
		key := *tag.Key
		value := *tag.Value
		tags[key] = value
	}

	return tags
}

func matchMaps(primary map[string]string, target map[string]string) bool {
	if len(primary) > len(target) {
		return false
	}

	for key, value := range primary {
		if targetValue, ok := target[key]; ok {
			if strings.EqualFold(value, targetValue) {
				continue
			}
		}

		return false
	}

	return true
}