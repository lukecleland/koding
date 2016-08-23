// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package ssm_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleSSM_AddTagsToResource() {
	svc := ssm.New(session.New())

	params := &ssm.AddTagsToResourceInput{
		ResourceId:   aws.String("ResourceId"),             // Required
		ResourceType: aws.String("ResourceTypeForTagging"), // Required
		Tags: []*ssm.Tag{ // Required
			{ // Required
				Key:   aws.String("TagKey"),   // Required
				Value: aws.String("TagValue"), // Required
			},
			// More values...
		},
	}
	resp, err := svc.AddTagsToResource(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSSM_CancelCommand() {
	svc := ssm.New(session.New())

	params := &ssm.CancelCommandInput{
		CommandId: aws.String("CommandId"), // Required
		InstanceIds: []*string{
			aws.String("InstanceId"), // Required
			// More values...
		},
	}
	resp, err := svc.CancelCommand(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSSM_CreateActivation() {
	svc := ssm.New(session.New())

	params := &ssm.CreateActivationInput{
		IamRole:             aws.String("IamRole"), // Required
		DefaultInstanceName: aws.String("DefaultInstanceName"),
		Description:         aws.String("ActivationDescription"),
		ExpirationDate:      aws.Time(time.Now()),
		RegistrationLimit:   aws.Int64(1),
	}
	resp, err := svc.CreateActivation(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSSM_CreateAssociation() {
	svc := ssm.New(session.New())

	params := &ssm.CreateAssociationInput{
		InstanceId: aws.String("InstanceId"),   // Required
		Name:       aws.String("DocumentName"), // Required
		Parameters: map[string][]*string{
			"Key": { // Required
				aws.String("ParameterValue"), // Required
				// More values...
			},
			// More values...
		},
	}
	resp, err := svc.CreateAssociation(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSSM_CreateAssociationBatch() {
	svc := ssm.New(session.New())

	params := &ssm.CreateAssociationBatchInput{
		Entries: []*ssm.CreateAssociationBatchRequestEntry{ // Required
			{ // Required
				InstanceId: aws.String("InstanceId"),
				Name:       aws.String("DocumentName"),
				Parameters: map[string][]*string{
					"Key": { // Required
						aws.String("ParameterValue"), // Required
						// More values...
					},
					// More values...
				},
			},
			// More values...
		},
	}
	resp, err := svc.CreateAssociationBatch(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSSM_CreateDocument() {
	svc := ssm.New(session.New())

	params := &ssm.CreateDocumentInput{
		Content: aws.String("DocumentContent"), // Required
		Name:    aws.String("DocumentName"),    // Required
	}
	resp, err := svc.CreateDocument(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSSM_DeleteActivation() {
	svc := ssm.New(session.New())

	params := &ssm.DeleteActivationInput{
		ActivationId: aws.String("ActivationId"), // Required
	}
	resp, err := svc.DeleteActivation(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSSM_DeleteAssociation() {
	svc := ssm.New(session.New())

	params := &ssm.DeleteAssociationInput{
		InstanceId: aws.String("InstanceId"),   // Required
		Name:       aws.String("DocumentName"), // Required
	}
	resp, err := svc.DeleteAssociation(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSSM_DeleteDocument() {
	svc := ssm.New(session.New())

	params := &ssm.DeleteDocumentInput{
		Name: aws.String("DocumentName"), // Required
	}
	resp, err := svc.DeleteDocument(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSSM_DeregisterManagedInstance() {
	svc := ssm.New(session.New())

	params := &ssm.DeregisterManagedInstanceInput{
		InstanceId: aws.String("ManagedInstanceId"), // Required
	}
	resp, err := svc.DeregisterManagedInstance(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSSM_DescribeActivations() {
	svc := ssm.New(session.New())

	params := &ssm.DescribeActivationsInput{
		Filters: []*ssm.DescribeActivationsFilter{
			{ // Required
				FilterKey: aws.String("DescribeActivationsFilterKeys"),
				FilterValues: []*string{
					aws.String("String"), // Required
					// More values...
				},
			},
			// More values...
		},
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("NextToken"),
	}
	resp, err := svc.DescribeActivations(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSSM_DescribeAssociation() {
	svc := ssm.New(session.New())

	params := &ssm.DescribeAssociationInput{
		InstanceId: aws.String("InstanceId"),   // Required
		Name:       aws.String("DocumentName"), // Required
	}
	resp, err := svc.DescribeAssociation(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSSM_DescribeDocument() {
	svc := ssm.New(session.New())

	params := &ssm.DescribeDocumentInput{
		Name: aws.String("DocumentARN"), // Required
	}
	resp, err := svc.DescribeDocument(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSSM_DescribeDocumentPermission() {
	svc := ssm.New(session.New())

	params := &ssm.DescribeDocumentPermissionInput{
		Name:           aws.String("DocumentName"),           // Required
		PermissionType: aws.String("DocumentPermissionType"), // Required
	}
	resp, err := svc.DescribeDocumentPermission(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSSM_DescribeInstanceInformation() {
	svc := ssm.New(session.New())

	params := &ssm.DescribeInstanceInformationInput{
		InstanceInformationFilterList: []*ssm.InstanceInformationFilter{
			{ // Required
				Key: aws.String("InstanceInformationFilterKey"), // Required
				ValueSet: []*string{ // Required
					aws.String("InstanceInformationFilterValue"), // Required
					// More values...
				},
			},
			// More values...
		},
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("NextToken"),
	}
	resp, err := svc.DescribeInstanceInformation(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSSM_GetDocument() {
	svc := ssm.New(session.New())

	params := &ssm.GetDocumentInput{
		Name: aws.String("DocumentARN"), // Required
	}
	resp, err := svc.GetDocument(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSSM_ListAssociations() {
	svc := ssm.New(session.New())

	params := &ssm.ListAssociationsInput{
		AssociationFilterList: []*ssm.AssociationFilter{ // Required
			{ // Required
				Key:   aws.String("AssociationFilterKey"),   // Required
				Value: aws.String("AssociationFilterValue"), // Required
			},
			// More values...
		},
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("NextToken"),
	}
	resp, err := svc.ListAssociations(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSSM_ListCommandInvocations() {
	svc := ssm.New(session.New())

	params := &ssm.ListCommandInvocationsInput{
		CommandId: aws.String("CommandId"),
		Details:   aws.Bool(true),
		Filters: []*ssm.CommandFilter{
			{ // Required
				Key:   aws.String("CommandFilterKey"),   // Required
				Value: aws.String("CommandFilterValue"), // Required
			},
			// More values...
		},
		InstanceId: aws.String("InstanceId"),
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("NextToken"),
	}
	resp, err := svc.ListCommandInvocations(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSSM_ListCommands() {
	svc := ssm.New(session.New())

	params := &ssm.ListCommandsInput{
		CommandId: aws.String("CommandId"),
		Filters: []*ssm.CommandFilter{
			{ // Required
				Key:   aws.String("CommandFilterKey"),   // Required
				Value: aws.String("CommandFilterValue"), // Required
			},
			// More values...
		},
		InstanceId: aws.String("InstanceId"),
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("NextToken"),
	}
	resp, err := svc.ListCommands(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSSM_ListDocuments() {
	svc := ssm.New(session.New())

	params := &ssm.ListDocumentsInput{
		DocumentFilterList: []*ssm.DocumentFilter{
			{ // Required
				Key:   aws.String("DocumentFilterKey"),   // Required
				Value: aws.String("DocumentFilterValue"), // Required
			},
			// More values...
		},
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("NextToken"),
	}
	resp, err := svc.ListDocuments(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSSM_ListTagsForResource() {
	svc := ssm.New(session.New())

	params := &ssm.ListTagsForResourceInput{
		ResourceId:   aws.String("ResourceId"),             // Required
		ResourceType: aws.String("ResourceTypeForTagging"), // Required
	}
	resp, err := svc.ListTagsForResource(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSSM_ModifyDocumentPermission() {
	svc := ssm.New(session.New())

	params := &ssm.ModifyDocumentPermissionInput{
		Name:           aws.String("DocumentName"),           // Required
		PermissionType: aws.String("DocumentPermissionType"), // Required
		AccountIdsToAdd: []*string{
			aws.String("AccountId"), // Required
			// More values...
		},
		AccountIdsToRemove: []*string{
			aws.String("AccountId"), // Required
			// More values...
		},
	}
	resp, err := svc.ModifyDocumentPermission(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSSM_RemoveTagsFromResource() {
	svc := ssm.New(session.New())

	params := &ssm.RemoveTagsFromResourceInput{
		ResourceId:   aws.String("ResourceId"),             // Required
		ResourceType: aws.String("ResourceTypeForTagging"), // Required
		TagKeys: []*string{ // Required
			aws.String("TagKey"), // Required
			// More values...
		},
	}
	resp, err := svc.RemoveTagsFromResource(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSSM_SendCommand() {
	svc := ssm.New(session.New())

	params := &ssm.SendCommandInput{
		DocumentName: aws.String("DocumentARN"), // Required
		InstanceIds: []*string{ // Required
			aws.String("InstanceId"), // Required
			// More values...
		},
		Comment:          aws.String("Comment"),
		DocumentHash:     aws.String("DocumentHash"),
		DocumentHashType: aws.String("DocumentHashType"),
		NotificationConfig: &ssm.NotificationConfig{
			NotificationArn: aws.String("NotificationArn"),
			NotificationEvents: []*string{
				aws.String("NotificationEvent"), // Required
				// More values...
			},
			NotificationType: aws.String("NotificationType"),
		},
		OutputS3BucketName: aws.String("S3BucketName"),
		OutputS3KeyPrefix:  aws.String("S3KeyPrefix"),
		Parameters: map[string][]*string{
			"Key": { // Required
				aws.String("ParameterValue"), // Required
				// More values...
			},
			// More values...
		},
		ServiceRoleArn: aws.String("ServiceRole"),
		TimeoutSeconds: aws.Int64(1),
	}
	resp, err := svc.SendCommand(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSSM_UpdateAssociationStatus() {
	svc := ssm.New(session.New())

	params := &ssm.UpdateAssociationStatusInput{
		AssociationStatus: &ssm.AssociationStatus{ // Required
			Date:           aws.Time(time.Now()),                // Required
			Message:        aws.String("StatusMessage"),         // Required
			Name:           aws.String("AssociationStatusName"), // Required
			AdditionalInfo: aws.String("StatusAdditionalInfo"),
		},
		InstanceId: aws.String("InstanceId"),   // Required
		Name:       aws.String("DocumentName"), // Required
	}
	resp, err := svc.UpdateAssociationStatus(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSSM_UpdateManagedInstanceRole() {
	svc := ssm.New(session.New())

	params := &ssm.UpdateManagedInstanceRoleInput{
		IamRole:    aws.String("IamRole"),           // Required
		InstanceId: aws.String("ManagedInstanceId"), // Required
	}
	resp, err := svc.UpdateManagedInstanceRole(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
