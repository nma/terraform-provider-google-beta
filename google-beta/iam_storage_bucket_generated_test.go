// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccStorageBucketIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   acctest.RandString(10),
		"role":            "roles/storage.objectViewer",
		"admin_role":      "roles/storage.admin",
		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccStorageBucketIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_storage_bucket_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("b/%s roles/storage.objectViewer", fmt.Sprintf("my-bucket%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccStorageBucketIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_storage_bucket_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("b/%s roles/storage.objectViewer", fmt.Sprintf("my-bucket%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccStorageBucketIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   acctest.RandString(10),
		"role":            "roles/storage.objectViewer",
		"admin_role":      "roles/storage.admin",
		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccStorageBucketIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_storage_bucket_iam_member.foo",
				ImportStateId:     fmt.Sprintf("b/%s roles/storage.objectViewer user:admin@hashicorptest.com", fmt.Sprintf("my-bucket%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccStorageBucketIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   acctest.RandString(10),
		"role":            "roles/storage.objectViewer",
		"admin_role":      "roles/storage.admin",
		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}
	context["service_account"] = getTestServiceAccountFromEnv(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccStorageBucketIamPolicy_basicGenerated(context),
			},
			{
				ResourceName:      "google_storage_bucket_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("b/%s", fmt.Sprintf("my-bucket%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccStorageBucketIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_storage_bucket_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("b/%s", fmt.Sprintf("my-bucket%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccStorageBucketIamBindingGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   acctest.RandString(10),
		"role":            "roles/storage.objectViewer",
		"admin_role":      "roles/storage.admin",
		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccStorageBucketIamBinding_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_storage_bucket_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("b/%s roles/storage.objectViewer %s", fmt.Sprintf("my-bucket%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccStorageBucketIamBindingGenerated_withAndWithoutCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   acctest.RandString(10),
		"role":            "roles/storage.objectViewer",
		"admin_role":      "roles/storage.admin",
		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccStorageBucketIamBinding_withAndWithoutConditionGenerated(context),
			},
			{
				ResourceName:      "google_storage_bucket_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("b/%s roles/storage.objectViewer", fmt.Sprintf("my-bucket%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_storage_bucket_iam_binding.foo2",
				ImportStateId:     fmt.Sprintf("b/%s roles/storage.objectViewer %s", fmt.Sprintf("my-bucket%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccStorageBucketIamMemberGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   acctest.RandString(10),
		"role":            "roles/storage.objectViewer",
		"admin_role":      "roles/storage.admin",
		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccStorageBucketIamMember_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_storage_bucket_iam_member.foo",
				ImportStateId:     fmt.Sprintf("b/%s roles/storage.objectViewer user:admin@hashicorptest.com %s", fmt.Sprintf("my-bucket%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccStorageBucketIamMemberGenerated_withAndWithoutCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   acctest.RandString(10),
		"role":            "roles/storage.objectViewer",
		"admin_role":      "roles/storage.admin",
		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccStorageBucketIamMember_withAndWithoutConditionGenerated(context),
			},
			{
				ResourceName:      "google_storage_bucket_iam_member.foo",
				ImportStateId:     fmt.Sprintf("b/%s roles/storage.objectViewer user:admin@hashicorptest.com", fmt.Sprintf("my-bucket%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_storage_bucket_iam_member.foo2",
				ImportStateId:     fmt.Sprintf("b/%s roles/storage.objectViewer user:admin@hashicorptest.com %s", fmt.Sprintf("my-bucket%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccStorageBucketIamPolicyGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   acctest.RandString(10),
		"role":            "roles/storage.objectViewer",
		"admin_role":      "roles/storage.admin",
		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}
	context["service_account"] = getTestServiceAccountFromEnv(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccStorageBucketIamPolicy_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_storage_bucket_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("b/%s", fmt.Sprintf("my-bucket%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccStorageBucketIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_bucket" "default" {
  name               = "tf-test-my-bucket%{random_suffix}"
  bucket_policy_only = true
}

resource "google_storage_bucket_iam_member" "foo" {
  bucket = google_storage_bucket.default.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccStorageBucketIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_bucket" "default" {
  name               = "tf-test-my-bucket%{random_suffix}"
  bucket_policy_only = true
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
  binding {
    role = "%{admin_role}"
    members = ["serviceAccount:%{service_account}"]
  }
}

resource "google_storage_bucket_iam_policy" "foo" {
  bucket = google_storage_bucket.default.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccStorageBucketIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_bucket" "default" {
  name               = "tf-test-my-bucket%{random_suffix}"
  bucket_policy_only = true
}

data "google_iam_policy" "foo" {
}

resource "google_storage_bucket_iam_policy" "foo" {
  bucket = google_storage_bucket.default.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccStorageBucketIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_bucket" "default" {
  name               = "tf-test-my-bucket%{random_suffix}"
  bucket_policy_only = true
}

resource "google_storage_bucket_iam_binding" "foo" {
  bucket = google_storage_bucket.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccStorageBucketIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_bucket" "default" {
  name               = "tf-test-my-bucket%{random_suffix}"
  bucket_policy_only = true
}

resource "google_storage_bucket_iam_binding" "foo" {
  bucket = google_storage_bucket.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:paddy@hashicorp.com"]
}
`, context)
}

func testAccStorageBucketIamBinding_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_bucket" "default" {
  name               = "tf-test-my-bucket%{random_suffix}"
  bucket_policy_only = true
}

resource "google_storage_bucket_iam_binding" "foo" {
  bucket = google_storage_bucket.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    title       = "%{condition_title}"
    description = "Expiring at midnight of 2019-12-31"
    expression  = "%{condition_expr}"
  }
}
`, context)
}

func testAccStorageBucketIamBinding_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_bucket" "default" {
  name               = "tf-test-my-bucket%{random_suffix}"
  bucket_policy_only = true
}

resource "google_storage_bucket_iam_binding" "foo" {
  bucket = google_storage_bucket.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}

resource "google_storage_bucket_iam_binding" "foo2" {
  bucket = google_storage_bucket.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    title       = "%{condition_title}"
    description = "Expiring at midnight of 2019-12-31"
    expression  = "%{condition_expr}"
  }
}
`, context)
}

func testAccStorageBucketIamMember_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_bucket" "default" {
  name               = "tf-test-my-bucket%{random_suffix}"
  bucket_policy_only = true
}

resource "google_storage_bucket_iam_member" "foo" {
  bucket = google_storage_bucket.default.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    title       = "%{condition_title}"
    description = "Expiring at midnight of 2019-12-31"
    expression  = "%{condition_expr}"
  }
}
`, context)
}

func testAccStorageBucketIamMember_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_bucket" "default" {
  name               = "tf-test-my-bucket%{random_suffix}"
  bucket_policy_only = true
}

resource "google_storage_bucket_iam_member" "foo" {
  bucket = google_storage_bucket.default.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}

resource "google_storage_bucket_iam_member" "foo2" {
  bucket = google_storage_bucket.default.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    title       = "%{condition_title}"
    description = "Expiring at midnight of 2019-12-31"
    expression  = "%{condition_expr}"
  }
}
`, context)
}

func testAccStorageBucketIamPolicy_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_bucket" "default" {
  name               = "tf-test-my-bucket%{random_suffix}"
  bucket_policy_only = true
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
    condition {
      title       = "%{condition_title}"
      description = "Expiring at midnight of 2019-12-31"
      expression  = "%{condition_expr}"
    }
  }
  binding {
    role = "%{admin_role}"
    members = ["serviceAccount:%{service_account}"]
  }
}

resource "google_storage_bucket_iam_policy" "foo" {
  bucket = google_storage_bucket.default.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}