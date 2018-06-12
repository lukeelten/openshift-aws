package aws

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"util"
	"configuration"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
)

func GetRegistryBucketName(config *configuration.InputVars) string {
	tags := make(map[string]string, 2)
	tags["ProjectId"] = config.ProjectId
	tags["Type"] = "registry"

	return getBucketNameByTags(tags)
}


func getBucketNameByTags(tags map[string]string) string {
	buckets := getAllBuckets()

	for _, bucket := range buckets {
		bucketTags := getBucketTags(bucket)

		if matchMaps(tags, bucketTags) {
			return bucket
		}
	}

	panic("Cannot find matching bucket")
}

func getAllBuckets() []string {
	filter := s3.ListBucketsInput{}

	result, err := S3Client.ListBuckets(&filter)
	util.ExitOnError("Error when getting bucket list via API", err)

	buckets := make([]string, len(result.Buckets))

	for i, bucket := range result.Buckets {
		buckets[i] = *bucket.Name
	}

	return buckets
}

func getBucketTags(bucket string) map[string]string {
	input := s3.GetBucketTaggingInput{Bucket: aws.String(bucket)}

	output, err := S3Client.GetBucketTagging(&input)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			if awsErr.Code() == "NoSuchTagSet" || awsErr.Code() == "NoSuchBucket" {
				// Error ok, just return empty map
				return make(map[string]string, 0)
			}
		}
	}

	util.ExitOnError("Cannot get tags for bucket:" + bucket, err)

	tags := make(map[string]string, len(output.TagSet))

	for _, tag := range output.TagSet {
		key := *tag.Key
		value := *tag.Value
		tags[key] = value
	}

	return tags
}