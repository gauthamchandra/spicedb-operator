package v1alpha1

import (
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func (c *SpiceDBCluster) NamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      c.Name,
		Namespace: c.Namespace,
	}
}

const (
	ConditionTypeValidating          = "Validating"
	ConditionValidatingFailed        = "ValidatingFailed"
	ConditionTypeMigrating           = "Migrating"
	ConditionTypeConfigWarnings      = "ConfigurationWarning"
	ConditionTypePreconditionsFailed = "PreconditionsFailed"
	ConditionTypeRolling             = "RollingDeployment"

	ConditionReasonMissingSecret = "MissingSecret"
)

func NewValidatingConfigCondition(secretHash string) metav1.Condition {
	return metav1.Condition{
		Type:               ConditionTypeValidating,
		Status:             metav1.ConditionTrue,
		Reason:             "ConfigChanged",
		LastTransitionTime: metav1.NewTime(time.Now()),
		Message:            fmt.Sprintf("Validating config with secret hash %q", secretHash),
	}
}

func NewInvalidConfigCondition(secretHash string, err error) metav1.Condition {
	return metav1.Condition{
		Type:               ConditionValidatingFailed,
		Status:             metav1.ConditionTrue,
		Reason:             "InvalidConfig",
		LastTransitionTime: metav1.NewTime(time.Now()),
		Message:            fmt.Sprintf("Error validating config with secret hash %q: %s", secretHash, err),
	}
}

func NewConfigWarningCondition(warning error) metav1.Condition {
	return metav1.Condition{
		Type:               ConditionTypeConfigWarnings,
		Status:             metav1.ConditionTrue,
		Reason:             "WarningsPresent",
		LastTransitionTime: metav1.NewTime(time.Now()),
		Message:            warning.Error(),
	}
}

func NewMigratingCondition(engine, headRevision string) metav1.Condition {
	return metav1.Condition{
		Type:               ConditionTypeMigrating,
		Status:             metav1.ConditionTrue,
		Reason:             "MigrationJobRunning",
		LastTransitionTime: metav1.NewTime(time.Now()),
		Message:            fmt.Sprintf("Migrating %s datastore to %s", engine, headRevision),
	}
}

func NewMigrationFailedCondition(engine, headRevision string, err error) metav1.Condition {
	return metav1.Condition{
		Type:               ConditionTypeMigrating,
		Status:             metav1.ConditionFalse,
		Reason:             "MigrationFailed",
		LastTransitionTime: metav1.NewTime(time.Now()),
		Message:            fmt.Sprintf("Migrating %s datastore to %s failed with error: %s", engine, headRevision, err),
	}
}

func NewMissingSecretCondition(nn types.NamespacedName) metav1.Condition {
	return metav1.Condition{
		Type:               ConditionTypePreconditionsFailed,
		Status:             metav1.ConditionTrue,
		Reason:             ConditionReasonMissingSecret,
		LastTransitionTime: metav1.NewTime(time.Now()),
		Message:            fmt.Sprintf("Secret %s not found", nn.String()),
	}
}

func NewRollingCondition(message string) metav1.Condition {
	return metav1.Condition{
		Type:               ConditionTypeRolling,
		Status:             metav1.ConditionTrue,
		Reason:             "WaitingForDeploymentAvailability",
		LastTransitionTime: metav1.NewTime(time.Now()),
		Message:            message,
	}
}
