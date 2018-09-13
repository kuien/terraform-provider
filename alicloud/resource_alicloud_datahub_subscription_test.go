package alicloud

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccAlicloudDatahubSubscription_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {

			testAccPreCheck(t)
		},

		// module name
		IDRefreshName: "alicloud_datahub_subscription.basic",
		Providers:     testAccProviders,
		CheckDestroy:  testAccCheckDatahubSubscriptionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDatahubSubscription,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDatahubProjectExist(
						"alicloud_datahub_project.basic"),
					testAccCheckDatahubTopicExist(
						"alicloud_datahub_topic.basic"),
					testAccCheckDatahubSubscriptionExist(
						"alicloud_datahub_subscription.basic"),
					resource.TestCheckResourceAttr(
						"alicloud_datahub_subscription.basic",
						"project_name", "tftestDatahubProject"),
					resource.TestCheckResourceAttr(
						"alicloud_datahub_subscription.basic",
						"topic_name", "tftestDatahubTopic"),
				),
			},
		},
	})
}

func testAccCheckDatahubSubscriptionExist(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found Datahub subscritpion: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Datahub Subscription ID is set")
		}

		dh := testAccProvider.Meta().(*AliyunClient).dhconn

		split := strings.Split(rs.Primary.ID, COLON_SEPARATED)
		projectName := split[0]
		topicName := split[1]
		subId := split[2]
		_, err := dh.GetSubscription(projectName, topicName, subId)

		if err != nil && !NotFoundError(err) {
			return err
		}
		return nil
	}
}

func testAccCheckDatahubSubscriptionDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "alicloud_datahub_subscription" {
			continue
		}

		dh := testAccProvider.Meta().(*AliyunClient).dhconn

		split := strings.Split(rs.Primary.ID, COLON_SEPARATED)
		projectName := split[0]
		topicName := split[1]
		subId := split[2]
		_, err := dh.GetSubscription(projectName, topicName, subId)

		if err != nil {
			if NotFoundError(err) {
				continue
			}
			return err
		}

		return fmt.Errorf("Datahub subscription %s still exists", rs.Primary.ID)
	}

	return nil
}

const testAccDatahubSubscription = `
provider "alicloud" {
    region = "cn-beijing"
}
variable "project_name" {
  default = "tftestDatahubProject"
}
variable "topic_name" {
  default = "tftestDatahubTopic"
}
variable "record_type" {
  default = "BLOB"
}
resource "alicloud_datahub_project" "basic" {
  name = "${var.project_name}"
  comment = "project for basic"
}
resource "alicloud_datahub_topic" "basic" {
  project_name = "${alicloud_datahub_project.basic.name}"
  topic_name = "${var.topic_name}"
  record_type = "${var.record_type}"
  shard_count = 3
  life_cycle = 7
  comment = "topic for basic."
}
resource "alicloud_datahub_subscription" "basic" {
  project_name = "${alicloud_datahub_project.basic.name}"
  topic_name = "${alicloud_datahub_topic.basic.topic_name}"
  comment = "subscription for basic."
}
`
