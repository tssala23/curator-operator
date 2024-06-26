/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ReportSpec defines the desired state of Report
type ReportSpec struct {
	//Namespace in which cron job is created
	CronjobNamespace string `json:"cronjobNamespace,omitempty"`

	//Schedule period for the Report Cronjob
	//ScheduleForReport string `json:"scheduleForReport,omitempty"`

	// Value for the Database Name Environment Variable
	DatabaseName string `json:"databaseName,omitempty"`

	//Value for the Database Password Environment Variable
	DatabasePassword string `json:"databasePassword,omitempty"`

	// Value for the Database User Environment Variable
	DatabaseUser string `json:"databaseUser,omitempty"`

	// Value for the Database HostName Environment Variable
	DatabaseHostName string `json:"databaseHostName,omitempty"`

	// Value for the Database Environment Variable in order to define the port which it should use. It will be used in its container as well
	DatabasePort string `json:"databasePort,omitempty"`

	// Frequency of the report generation
	ReportFrequency string `json:"reportFrequency,omitempty"`

	//If email has to be sent then this value should be set to true
	EmailReports bool `json:"emailReports,omitempty"`

	//Sender email address which will be used as a mail address
	EmailUser string `json:"emailUser,omitempty"`

	//Sender email password which does not need 2FA
	EmailPassword string `json:"emailPassword,omitempty"`

	//Recepients email who will receive the reports, should be in the format {"user1@gmail.com":{"cc":["cc1@gmail.com"]}}
	EmailRecepients string `json:"emailRecepients,omitempty"`
}

// ReportStatus defines the observed state of Report
type ReportStatus struct {
	//Name of the CronJob object created and managed by it
	CronJobName string `json:"cronJobName"`

	//CronJobStatus represents the current state of a cronjob
	CronJobStatus batchv1.CronJobStatus `json:"cronJobStatus"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Report is the Schema for the reports API
type Report struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ReportSpec   `json:"spec,omitempty"`
	Status ReportStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ReportList contains a list of Report
type ReportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Report `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Report{}, &ReportList{})
}
