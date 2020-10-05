package aws

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lexmodelbuildingservice"
	"github.com/hashicorp/aws-sdk-go-base/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccAwsLexBotAlias_basic(t *testing.T) {
	var v lexmodelbuildingservice.GetBotAliasOutput
	rName := "aws_lex_bot_alias.test"
	testBotAliasID := "test_bot_alias" + acctest.RandStringFromCharSet(8, acctest.CharSetAlpha)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAwsLexBotAliasDestroy(testBotAliasID, testBotAliasID),
		Steps: []resource.TestStep{
			{
				Config: composeConfig(
					testAccAwsLexBotConfig_intent(testBotAliasID),
					testAccAwsLexBotConfig_basic(testBotAliasID),
					testAccAwsLexBotAliasConfig_basic(testBotAliasID),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAwsLexBotAliasExists(rName, &v),

					resource.TestCheckResourceAttrSet(rName, "arn"),
					resource.TestCheckResourceAttrSet(rName, "checksum"),
					testAccCheckResourceAttrRfc3339(rName, "created_date"),
					resource.TestCheckResourceAttr(rName, "description", "Testing lex bot alias create."),
					testAccCheckResourceAttrRfc3339(rName, "last_updated_date"),
					resource.TestCheckResourceAttr(rName, "bot_name", testBotAliasID),
					resource.TestCheckResourceAttr(rName, "bot_version", "$LATEST"),
					resource.TestCheckResourceAttr(rName, "name", testBotAliasID),
				),
			},
			{
				ResourceName:      rName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAwsLexBotAlias_botVersion(t *testing.T) {
	var v lexmodelbuildingservice.GetBotAliasOutput
	rName := "aws_lex_bot_alias.test"
	testBotAliasID := "test_bot_alias" + acctest.RandStringFromCharSet(8, acctest.CharSetAlpha)

	// If this test runs in parallel with other Lex Bot tests, it loses its description
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAwsLexBotAliasDestroy(testBotAliasID, testBotAliasID),
		Steps: []resource.TestStep{
			{
				Config: composeConfig(
					testAccAwsLexBotConfig_intent(testBotAliasID),
					testAccAwsLexBotConfig_basic(testBotAliasID),
					testAccAwsLexBotAliasConfig_basic(testBotAliasID),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAwsLexBotAliasExists(rName, &v),
				),
			},
			{
				ResourceName:      rName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: composeConfig(
					testAccAwsLexBotConfig_intent(testBotAliasID),
					testAccAwsLexBotConfig_createVersion(testBotAliasID),
					testAccAwsLexBotAliasConfig_botVersionUpdate(testBotAliasID),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAwsLexBotAliasExists(rName, &v),
					resource.TestCheckResourceAttr(rName, "bot_version", "1"),
				),
			},
			{
				ResourceName:      rName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAwsLexBotAlias_conversationLogsText(t *testing.T) {
	var v lexmodelbuildingservice.GetBotAliasOutput
	rName := "aws_lex_bot_alias.test"
	testBotAliasID := "test_bot_alias" + acctest.RandStringFromCharSet(8, acctest.CharSetAlpha)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAwsLexBotAliasDestroy(testBotAliasID, testBotAliasID),
		Steps: []resource.TestStep{
			{
				Config: composeConfig(
					testAccAwsLexBotAliasConfig_conversationLogsTextSetup(testBotAliasID),
				),
			},
			{
				Config: composeConfig(
					testAccAwsLexBotConfig_intent(testBotAliasID),
					testAccAwsLexBotConfig_basic(testBotAliasID),
					testAccAwsLexBotAliasConfig_conversationLogsTextSetup(testBotAliasID),
					testAccAwsLexBotAliasConfig_conversationLogsText(testBotAliasID),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAwsLexBotAliasExists(rName, &v),
				),
			},
			{
				ResourceName:      rName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAwsLexBotAlias_conversationLogsAudio(t *testing.T) {
	var v lexmodelbuildingservice.GetBotAliasOutput
	rName := "aws_lex_bot_alias.test"
	testBotAliasID := acctest.RandStringFromCharSet(8, acctest.CharSetAlpha)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAwsLexBotAliasDestroy(testBotAliasID, testBotAliasID),
		Steps: []resource.TestStep{
			{
				Config: composeConfig(
					testAccAwsLexBotAliasConfig_conversationLogsAudioSetup(testBotAliasID),
				),
			},
			{
				Config: composeConfig(
					testAccAwsLexBotConfig_intent(testBotAliasID),
					testAccAwsLexBotConfig_basic(testBotAliasID),
					testAccAwsLexBotAliasConfig_conversationLogsAudioSetup(testBotAliasID),
					testAccAwsLexBotAliasConfig_conversationLogsAudio(testBotAliasID),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAwsLexBotAliasExists(rName, &v),
				),
			},
			{
				ResourceName:      rName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAwsLexBotAlias_description(t *testing.T) {
	var v lexmodelbuildingservice.GetBotAliasOutput
	rName := "aws_lex_bot_alias.test"
	testBotAliasID := "test_bot_alias" + acctest.RandStringFromCharSet(8, acctest.CharSetAlpha)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAwsLexBotAliasDestroy(testBotAliasID, testBotAliasID),
		Steps: []resource.TestStep{
			{
				Config: composeConfig(
					testAccAwsLexBotConfig_intent(testBotAliasID),
					testAccAwsLexBotConfig_basic(testBotAliasID),
					testAccAwsLexBotAliasConfig_basic(testBotAliasID),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAwsLexBotAliasExists(rName, &v),
				),
			},
			{
				ResourceName:      rName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: composeConfig(
					testAccAwsLexBotConfig_intent(testBotAliasID),
					testAccAwsLexBotConfig_basic(testBotAliasID),
					testAccAwsLexBotAliasConfig_descriptionUpdate(testBotAliasID),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAwsLexBotAliasExists(rName, &v),
					resource.TestCheckResourceAttr(rName, "description", "Testing lex bot alias update."),
				),
			},
			{
				ResourceName:      rName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAwsLexBotAlias_disappears(t *testing.T) {
	var v lexmodelbuildingservice.GetBotAliasOutput
	rName := "aws_lex_bot_alias.test"
	testBotAliasID := "test_bot_alias" + acctest.RandStringFromCharSet(8, acctest.CharSetAlpha)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAwsLexBotDestroy,
		Steps: []resource.TestStep{
			{
				Config: composeConfig(
					testAccAwsLexBotConfig_intent(testBotAliasID),
					testAccAwsLexBotConfig_basic(testBotAliasID),
					testAccAwsLexBotAliasConfig_basic(testBotAliasID),
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsLexBotAliasExists(rName, &v),
					testAccCheckResourceDisappears(testAccProvider, resourceAwsLexBotAlias(), rName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckAwsLexBotAliasExists(rName string, output *lexmodelbuildingservice.GetBotAliasOutput) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[rName]
		if !ok {
			return fmt.Errorf("Not found: %s", rName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Lex bot alias ID is set")
		}

		botName := rs.Primary.Attributes["bot_name"]
		botAliasName := rs.Primary.Attributes["name"]

		var err error
		conn := testAccProvider.Meta().(*AWSClient).lexmodelconn

		output, err = conn.GetBotAlias(&lexmodelbuildingservice.GetBotAliasInput{
			BotName: aws.String(botName),
			Name:    aws.String(botAliasName),
		})
		if tfawserr.ErrCodeEquals(err, lexmodelbuildingservice.ErrCodeNotFoundException) {
			return fmt.Errorf("error bot alias '%q' not found", rs.Primary.ID)
		}
		if err != nil {
			return fmt.Errorf("error getting bot alias '%q': %w", rs.Primary.ID, err)
		}

		return nil
	}
}

func testAccCheckAwsLexBotAliasDestroy(botName, botAliasName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*AWSClient).lexmodelconn

		_, err := conn.GetBotAlias(&lexmodelbuildingservice.GetBotAliasInput{
			BotName: aws.String(botName),
			Name:    aws.String(botAliasName),
		})

		if err != nil {
			if isAWSErr(err, lexmodelbuildingservice.ErrCodeNotFoundException, "") {
				return nil
			}

			return fmt.Errorf("error getting bot alias '%s': %s", botAliasName, err)
		}

		return fmt.Errorf("error bot alias still exists after delete, %s", botAliasName)
	}
}

func testAccAwsLexBotAliasConfig_basic(rName string) string {
	return fmt.Sprintf(`
resource "aws_lex_bot_alias" "test" {
  bot_name    = aws_lex_bot.test.name
  bot_version = aws_lex_bot.test.version
  description = "Testing lex bot alias create."
  name        = "%s"
}
`, rName)
}

func testAccAwsLexBotAliasConfig_botVersionUpdate(rName string) string {
	return fmt.Sprintf(`
resource "aws_lex_bot_alias" "test" {
  bot_name    = aws_lex_bot.test.name
  bot_version = "1"
  description = "Testing lex bot alias create."
  name        = "%s"
}
`, rName)
}

func testAccAwsLexBotAliasConfig_conversationLogsTextSetup(rName string) string {
	return fmt.Sprintf(`
data "aws_iam_policy_document" "lex_assume_role_policy" {
  statement {
    effect  = "Allow"
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["lex.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "test" {
  name               = "%[1]s"
  assume_role_policy = data.aws_iam_policy_document.lex_assume_role_policy.json
}

data "aws_iam_policy_document" "lex_cloud_watch_logs_policy" {
  statement {
    effect = "Allow"
    actions = [
      "logs:CreateLogStream",
      "logs:PutLogEvents",
    ]
    resources = [
      aws_cloudwatch_log_group.test.arn,
    ]
  }
}

resource "aws_iam_role_policy" "test" {
  name   = "%[1]s"
  role   = aws_iam_role.test.id
  policy = data.aws_iam_policy_document.lex_cloud_watch_logs_policy.json
}

resource "aws_cloudwatch_log_group" "test" {
  name = "%[1]s"
}
`, rName)
}

func testAccAwsLexBotAliasConfig_conversationLogsText(rName string) string {
	return fmt.Sprintf(`
resource "aws_lex_bot_alias" "test" {
  bot_name    = aws_lex_bot.test.name
  bot_version = aws_lex_bot.test.version
  description = "Testing lex bot alias create."
  name        = "%[1]s"
  conversation_logs {
    iam_role_arn = aws_iam_role.test.arn
    log_settings {
      destination  = "CLOUDWATCH_LOGS"
      log_type     = "TEXT"
      resource_arn = aws_cloudwatch_log_group.test.arn
    }
  }
}
`, rName)
}

func testAccAwsLexBotAliasConfig_conversationLogsAudioSetup(rName string) string {
	return fmt.Sprintf(`
data "aws_iam_policy_document" "lex_assume_role_policy" {
  statement {
    effect  = "Allow"
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["lex.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "test" {
  name               = "%[1]s"
  assume_role_policy = data.aws_iam_policy_document.lex_assume_role_policy.json
}

data "aws_iam_policy_document" "lex_s3_policy" {
  statement {
    effect = "Allow"
    actions = [
      "s3:PutObject",
    ]
    resources = [
      aws_s3_bucket.test.arn,
    ]
  }
}

resource "aws_iam_role_policy" "test" {
  name   = "%[1]s"
  role   = aws_iam_role.test.id
  policy = data.aws_iam_policy_document.lex_s3_policy.json
}

resource "aws_s3_bucket" "test" {
  bucket = "%[1]s"
}
`, rName)
}

func testAccAwsLexBotAliasConfig_conversationLogsAudio(rName string) string {
	return fmt.Sprintf(`
resource "aws_lex_bot_alias" "test" {
  bot_name    = aws_lex_bot.test.name
  bot_version = aws_lex_bot.test.version
  description = "Testing lex bot alias create."
  name        = "%[1]s"
  conversation_logs {
    iam_role_arn = aws_iam_role.test.arn
    log_settings {
      destination  = "S3"
      log_type     = "AUDIO"
      resource_arn = aws_s3_bucket.test.arn
    }
  }
}
`, rName)
}

func testAccAwsLexBotAliasConfig_descriptionUpdate(rName string) string {
	return fmt.Sprintf(`
resource "aws_lex_bot_alias" "test" {
  bot_name    = aws_lex_bot.test.name
  bot_version = aws_lex_bot.test.version
  description = "Testing lex bot alias update."
  name        = "%s"
}
`, rName)
}
