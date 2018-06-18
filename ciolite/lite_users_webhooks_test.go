package ciolite

import (
	"encoding/json"
	"reflect"
	"testing"
)

// TestReceivingWebhookEmptyAddresses tests receiving and parsing WebhookCallback with empty addresses field
func TestReceivingWebhookEmptyAddresses(t *testing.T) {
	t.Parallel()

	emptyAddressesJSON := `
{
	"account_id": "abc4567XYZ",
	"webhook_id": "abc4567XYZ",
	"timestamp": 1467254577,
	"message_data": {
		"message_id": "aaaaa111aaaa11aa",
		"addresses": [],
		"references": [],
		"flags": {
			"flagged": false,
			"answered": false,
			"draft": false,
			"seen": false
		},
		"sources": [],
		"email_accounts": [],
		"files": [],
		"person_info": [],
		"date": 1464742490,
		"subject": "Test Subject",
		"folders": ["Inbox"]
	},
	"user_id": "abc4567XYZ",
	"token": "fake578Token",
	"signature": "fake123Signature"
}`

	// Test Empty Addresses Array
	var emptyAddresses WebhookCallback
	Must(json.Unmarshal([]byte(emptyAddressesJSON), &emptyAddresses))

	if emptyAddresses.AccountID != "abc4567XYZ" {
		t.Error("Expected AccountID: ", "abc4567XYZ", "; Got: ", emptyAddresses.AccountID)
	}

	if len(emptyAddresses.MessageData.Folders) != 1 || emptyAddresses.MessageData.Folders[0] != "Inbox" {
		t.Error("Expected MessageData.Folders: ", "[Inbox]", "; Got: ", emptyAddresses.MessageData.Folders)
	}

	emptyExpected := WebhookMessageDataAddresses{}
	if !reflect.DeepEqual(emptyAddresses.MessageData.Addresses, emptyExpected) {
		t.Error("Expected MessageData.Addresses: ", emptyExpected, "; Got: ", emptyAddresses.MessageData.Addresses)
	}
}

// TestSimulatedReceivingWebhookFullAddresses tests receiving and parsing WebhookCallback with filled in addresses field
func TestSimulatedReceivingWebhookFullAddresses(t *testing.T) {
	t.Parallel()

	fullAddressesJSON := `
{
	"account_id": "abc4567XYZ",
	"webhook_id": "abc4567XYZ",
	"token": "fake578Token",
	"signature": "fake123Signature",
	"timestamp": 1467254202,
	"message_data": {
		"message_id": "aaaaa111aaaa11aa",
		"subject": "Test Subject",
		"folders": ["Inbox"],
		"date": 1446919784,
		"date_received": 1446919783,
		"addresses": {
			"from": {
				"email": "from@test.com",
				"name": "John"
			}
		},
		"person_info": {
			"noreply@test.com": {
				"thumbnail": "https://test.com/test.png"
			}
		},
		"flags": {
			"seen": true
		},
		"email_accounts": []
	}
}`

	// Test filled addresses field
	var fullAddresses WebhookCallback
	Must(json.Unmarshal([]byte(fullAddressesJSON), &fullAddresses))

	if fullAddresses.AccountID != "abc4567XYZ" {
		t.Error("Expected AccountID: ", "abc4567XYZ", "; Got: ", fullAddresses.AccountID)
	}

	if len(fullAddresses.MessageData.Folders) != 1 || fullAddresses.MessageData.Folders[0] != "Inbox" {
		t.Error("Expected MessageData.Folders: ", "[Inbox]", "; Got: ", fullAddresses.MessageData.Folders)
	}

	addressesExpected := WebhookMessageDataAddresses{
		From: Address{
			Email: "from@test.com",
			Name:  "John",
		},
	}
	if !reflect.DeepEqual(fullAddresses.MessageData.Addresses, addressesExpected) {
		t.Error("Expected MessageData.Addresses: ", addressesExpected, "; Got: ", fullAddresses.MessageData.Addresses)
	}

	personInfoExpected := map[string]map[string]string{"noreply@test.com": {"thumbnail": "https://test.com/test.png"}}

	if !reflect.DeepEqual(fullAddresses.MessageData.PersonInfo["noreply@test.com"], personInfoExpected["noreply@test.com"]) {
		t.Error("Expected MessageData.PersonInfo: ", personInfoExpected, "; Got: ", fullAddresses.MessageData.PersonInfo)
	}
}
