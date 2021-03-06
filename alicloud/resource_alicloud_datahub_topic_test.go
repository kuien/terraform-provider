package alicloud

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccAlicloudDatahubTopic_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},

		// module name
		IDRefreshName: "alicloud_datahub_topic.basic",
		Providers:     testAccProviders,
		CheckDestroy:  testAccCheckDatahubTopicDestroy,

		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDatahubTopic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDatahubProjectExist(
						"alicloud_datahub_project.basic"),
					testAccCheckDatahubTopicExist(
						"alicloud_datahub_topic.basic"),
					resource.TestCheckResourceAttr(
						"alicloud_datahub_topic.basic",
						"topic_name", "tftestDatahubTopicBasic"),
				),
			},
		},
	})
}

func TestAccAlicloudDatahubTopic_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},

		// module name
		IDRefreshName: "alicloud_datahub_topic.basic",
		Providers:     testAccProviders,
		CheckDestroy:  testAccCheckDatahubTopicDestroy,

		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDatahubTopic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDatahubTopicExist(
						"alicloud_datahub_topic.basic"),
					resource.TestCheckResourceAttr(
						"alicloud_datahub_topic.basic",
						"life_cycle", "7"),
				),
			},

			resource.TestStep{
				Config: testAccDatahubTopicUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDatahubTopicExist(
						"alicloud_datahub_topic.basic"),
					resource.TestCheckResourceAttr(
						"alicloud_datahub_topic.basic",
						"life_cycle", "1"),
				),
			},
		},
	})
}

func testAccCheckDatahubTopicExist(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found Datahub topic: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Datahub topic ID is set")
		}

		dh := testAccProvider.Meta().(*AliyunClient).dhconn

		split := strings.Split(rs.Primary.ID, COLON_SEPARATED)
		projectName := split[0]
		topicName := split[1]
		_, err := dh.GetTopic(topicName, projectName)

		// XXX DEBUG only
		// topic, err := dh.GetTopic(topicName, projectName)
		// fmt.Printf("\nXXX:life_cycle:%d\n", topic.Lifecycle)
		// fmt.Printf("XXX:comment:%s\n", topic.Comment)
		// fmt.Printf("XXX:create_time:%s\n", convUint64ToDate(topic.CreateTime))
		// fmt.Printf("XXX:last_modify_time:%s\n", convUint64ToDate(topic.LastModifyTime))

		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckDatahubTopicDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "alicloud_datahub_topic" {
			continue
		}

		dh := testAccProvider.Meta().(*AliyunClient).dhconn

		split := strings.Split(rs.Primary.ID, COLON_SEPARATED)
		projectName := split[0]
		topicName := split[1]
		_, err := dh.GetTopic(topicName, projectName)

		if err != nil {
			continue
		}

		return fmt.Errorf("Datahub topic %s still exists", rs.Primary.ID)
	}

	return nil
}

const testAccDatahubTopic = `
provider "alicloud" {
    region = "cn-beijing"
}
variable "project_name" {
  default = "tftestDatahubProject"
}
variable "topic_name" {
  default = "tftestDatahubTopicBasic"
}
resource "alicloud_datahub_project" "basic" {
  name = "${var.project_name}"
  comment = "Datahub project ${var.project_name} is used for terraform test only."
}
resource "alicloud_datahub_topic" "basic" {
  project_name = "${var.project_name}"
  topic_name = "${var.topic_name}"
  shard_count = 3
  life_cycle = 7
  comment = "Datahub topic ${var.topic_name} is used for terraform test only."
}
`
const testAccDatahubTopicUpdate = `
provider "alicloud" {
    region = "cn-beijing"
}
variable "project_name" {
  default = "tftestDatahubProject"
}
variable "topic_name" {
  default = "tftestDatahubTopicBasic"
}
resource "alicloud_datahub_project" "basic" {
  name = "${var.project_name}"
  comment = "Datahub project ${var.project_name} is used for terraform test only."
}
resource "alicloud_datahub_topic" "basic" {
  project_name = "${var.project_name}"
  topic_name = "${var.topic_name}"
  shard_count = 3
  life_cycle = 1
  comment = "Datahub topic ${var.topic_name} is used for terraform test only.\nNow being updated."
}
`
