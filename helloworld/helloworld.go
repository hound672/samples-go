package helloworld

import (
	"context"
	"fmt"
	"go.temporal.io/sdk/temporal"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"

	// TODO(cretz): Remove when tagged
	_ "go.temporal.io/sdk/contrib/tools/workflowcheck/determinism"
)

// Workflow is a Hello World workflow definition.
func Workflow(ctx workflow.Context, name string) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 3,
		},
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("HelloWorld workflow started", "name", name)

	var result string
	err := workflow.ExecuteActivity(ctx, "Activity", name).Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return "", err
	}

	if err := workflow.ExecuteActivity(ctx, Activity2, name).Get(ctx, &result); err != nil {
		logger.Error("222222 failed", "Error", err)
		return "", err
	}

	logger.Info("HelloWorld workflow completed.", "result", result)

	return result + "SOME PAYLOAD", nil
}

var cnt int

func Activity(ctx context.Context, name string) (string, error) {
	cnt++
	fmt.Printf("ACTIVITY!!!!!!!!!!!!!!!!: %d\n", cnt)
	//if cnt%2 == 0 {
	//if cnt > 2 {
	//	fmt.Printf("FAIL!!!!!!!!!!!!!!!!!\n")
	//	return "", fmt.Errorf("SOME ERROR HAPPENED")
	//}

	logger := activity.GetLogger(ctx)
	logger.Info("Activity", "name", name)
	return "Hello " + name + "!2222", nil
}

func Activity2(ctx context.Context, name string) (string, error) {
	fmt.Printf("22222222222222 ACTIVITY!!!!!!!!!!!!!!!!: %d\n", cnt)

	logger := activity.GetLogger(ctx)
	logger.Info("Activity", "name", name)
	return "ACT 2 Hello " + name + "!2222", nil
}
