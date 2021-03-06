// Code generated by protoc-gen-go.
// source: config.proto
// DO NOT EDIT!

/*
Package io_prometheus_alertmanager is a generated protocol buffer package.

It is generated from these files:
	config.proto

It has these top-level messages:
	PagerDutyConfig
	EmailConfig
	PushoverConfig
	HipChatConfig
	SlackConfig
	FlowdockConfig
	NotificationConfig
	Filter
	AggregationRule
	InhibitRule
	AlertManagerConfig
*/
package io_prometheus_alertmanager

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

// Configuration for notification via PagerDuty.
type PagerDutyConfig struct {
	// PagerDuty service key, see:
	// http://developer.pagerduty.com/documentation/integration/events
	ServiceKey       *string `protobuf:"bytes,1,opt,name=service_key" json:"service_key,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PagerDutyConfig) Reset()         { *m = PagerDutyConfig{} }
func (m *PagerDutyConfig) String() string { return proto.CompactTextString(m) }
func (*PagerDutyConfig) ProtoMessage()    {}

func (m *PagerDutyConfig) GetServiceKey() string {
	if m != nil && m.ServiceKey != nil {
		return *m.ServiceKey
	}
	return ""
}

// Configuration for notification via mail.
type EmailConfig struct {
	// Email address to notify.
	Email *string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	// Notify when resolved.
	SendResolved     *bool  `protobuf:"varint,2,opt,name=send_resolved,def=0" json:"send_resolved,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *EmailConfig) Reset()         { *m = EmailConfig{} }
func (m *EmailConfig) String() string { return proto.CompactTextString(m) }
func (*EmailConfig) ProtoMessage()    {}

const Default_EmailConfig_SendResolved bool = false

func (m *EmailConfig) GetEmail() string {
	if m != nil && m.Email != nil {
		return *m.Email
	}
	return ""
}

func (m *EmailConfig) GetSendResolved() bool {
	if m != nil && m.SendResolved != nil {
		return *m.SendResolved
	}
	return Default_EmailConfig_SendResolved
}

// Configuration for notification via pushover.net.
type PushoverConfig struct {
	// Pushover token.
	Token *string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	// Pushover user_key.
	UserKey *string `protobuf:"bytes,2,opt,name=user_key" json:"user_key,omitempty"`
	// Notify when resolved.
	SendResolved     *bool  `protobuf:"varint,3,opt,name=send_resolved,def=0" json:"send_resolved,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *PushoverConfig) Reset()         { *m = PushoverConfig{} }
func (m *PushoverConfig) String() string { return proto.CompactTextString(m) }
func (*PushoverConfig) ProtoMessage()    {}

const Default_PushoverConfig_SendResolved bool = false

func (m *PushoverConfig) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

func (m *PushoverConfig) GetUserKey() string {
	if m != nil && m.UserKey != nil {
		return *m.UserKey
	}
	return ""
}

func (m *PushoverConfig) GetSendResolved() bool {
	if m != nil && m.SendResolved != nil {
		return *m.SendResolved
	}
	return Default_PushoverConfig_SendResolved
}

// Configuration for notification via HipChat.
type HipChatConfig struct {
	// HipChat auth token, (https://www.hipchat.com/docs/api/auth).
	AuthToken *string `protobuf:"bytes,1,opt,name=auth_token" json:"auth_token,omitempty"`
	// HipChat room id, (https://www.hipchat.com/rooms/ids).
	RoomId *int32 `protobuf:"varint,2,opt,name=room_id" json:"room_id,omitempty"`
	// Color of message when triggered.
	Color *string `protobuf:"bytes,3,opt,name=color,def=purple" json:"color,omitempty"`
	// Color of message when resolved.
	ColorResolved *string `protobuf:"bytes,5,opt,name=color_resolved,def=green" json:"color_resolved,omitempty"`
	// Should this message notify or not.
	Notify *bool `protobuf:"varint,4,opt,name=notify,def=0" json:"notify,omitempty"`
	// Notify when resolved.
	SendResolved     *bool  `protobuf:"varint,6,opt,name=send_resolved,def=0" json:"send_resolved,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *HipChatConfig) Reset()         { *m = HipChatConfig{} }
func (m *HipChatConfig) String() string { return proto.CompactTextString(m) }
func (*HipChatConfig) ProtoMessage()    {}

const Default_HipChatConfig_Color string = "purple"
const Default_HipChatConfig_ColorResolved string = "green"
const Default_HipChatConfig_Notify bool = false
const Default_HipChatConfig_SendResolved bool = false

func (m *HipChatConfig) GetAuthToken() string {
	if m != nil && m.AuthToken != nil {
		return *m.AuthToken
	}
	return ""
}

func (m *HipChatConfig) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *HipChatConfig) GetColor() string {
	if m != nil && m.Color != nil {
		return *m.Color
	}
	return Default_HipChatConfig_Color
}

func (m *HipChatConfig) GetColorResolved() string {
	if m != nil && m.ColorResolved != nil {
		return *m.ColorResolved
	}
	return Default_HipChatConfig_ColorResolved
}

func (m *HipChatConfig) GetNotify() bool {
	if m != nil && m.Notify != nil {
		return *m.Notify
	}
	return Default_HipChatConfig_Notify
}

func (m *HipChatConfig) GetSendResolved() bool {
	if m != nil && m.SendResolved != nil {
		return *m.SendResolved
	}
	return Default_HipChatConfig_SendResolved
}

// Configuration for notification via Slack.
type SlackConfig struct {
	// Slack webhook url, (https://api.slack.com/incoming-webhooks).
	WebhookUrl *string `protobuf:"bytes,1,opt,name=webhook_url" json:"webhook_url,omitempty"`
	// Slack channel override, (like #other-channel or @username).
	Channel *string `protobuf:"bytes,2,opt,name=channel" json:"channel,omitempty"`
	// Color of message when triggered.
	Color *string `protobuf:"bytes,3,opt,name=color,def=warning" json:"color,omitempty"`
	// Color of message when resolved.
	ColorResolved *string `protobuf:"bytes,4,opt,name=color_resolved,def=good" json:"color_resolved,omitempty"`
	// Notify when resolved.
	SendResolved     *bool  `protobuf:"varint,5,opt,name=send_resolved,def=0" json:"send_resolved,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SlackConfig) Reset()         { *m = SlackConfig{} }
func (m *SlackConfig) String() string { return proto.CompactTextString(m) }
func (*SlackConfig) ProtoMessage()    {}

const Default_SlackConfig_Color string = "warning"
const Default_SlackConfig_ColorResolved string = "good"
const Default_SlackConfig_SendResolved bool = false

func (m *SlackConfig) GetWebhookUrl() string {
	if m != nil && m.WebhookUrl != nil {
		return *m.WebhookUrl
	}
	return ""
}

func (m *SlackConfig) GetChannel() string {
	if m != nil && m.Channel != nil {
		return *m.Channel
	}
	return ""
}

func (m *SlackConfig) GetColor() string {
	if m != nil && m.Color != nil {
		return *m.Color
	}
	return Default_SlackConfig_Color
}

func (m *SlackConfig) GetColorResolved() string {
	if m != nil && m.ColorResolved != nil {
		return *m.ColorResolved
	}
	return Default_SlackConfig_ColorResolved
}

func (m *SlackConfig) GetSendResolved() bool {
	if m != nil && m.SendResolved != nil {
		return *m.SendResolved
	}
	return Default_SlackConfig_SendResolved
}

type FlowdockConfig struct {
	// Flowdock flow API token.
	ApiToken *string `protobuf:"bytes,1,opt,name=api_token" json:"api_token,omitempty"`
	// Flowdock from_address.
	FromAddress *string `protobuf:"bytes,2,opt,name=from_address" json:"from_address,omitempty"`
	// Flowdock flow tags.
	Tag []string `protobuf:"bytes,3,rep,name=tag" json:"tag,omitempty"`
	// Notify when resolved.
	SendResolved     *bool  `protobuf:"varint,4,opt,name=send_resolved,def=0" json:"send_resolved,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *FlowdockConfig) Reset()         { *m = FlowdockConfig{} }
func (m *FlowdockConfig) String() string { return proto.CompactTextString(m) }
func (*FlowdockConfig) ProtoMessage()    {}

const Default_FlowdockConfig_SendResolved bool = false

func (m *FlowdockConfig) GetApiToken() string {
	if m != nil && m.ApiToken != nil {
		return *m.ApiToken
	}
	return ""
}

func (m *FlowdockConfig) GetFromAddress() string {
	if m != nil && m.FromAddress != nil {
		return *m.FromAddress
	}
	return ""
}

func (m *FlowdockConfig) GetTag() []string {
	if m != nil {
		return m.Tag
	}
	return nil
}

func (m *FlowdockConfig) GetSendResolved() bool {
	if m != nil && m.SendResolved != nil {
		return *m.SendResolved
	}
	return Default_FlowdockConfig_SendResolved
}

// Notification configuration definition.
type NotificationConfig struct {
	// Name of this NotificationConfig. Referenced from AggregationRule.
	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Zero or more PagerDuty notification configurations.
	PagerdutyConfig []*PagerDutyConfig `protobuf:"bytes,2,rep,name=pagerduty_config" json:"pagerduty_config,omitempty"`
	// Zero or more email notification configurations.
	EmailConfig []*EmailConfig `protobuf:"bytes,3,rep,name=email_config" json:"email_config,omitempty"`
	// Zero or more pushover notification configurations.
	PushoverConfig []*PushoverConfig `protobuf:"bytes,4,rep,name=pushover_config" json:"pushover_config,omitempty"`
	// Zero or more hipchat notification configurations.
	HipchatConfig []*HipChatConfig `protobuf:"bytes,5,rep,name=hipchat_config" json:"hipchat_config,omitempty"`
	// Zero or more slack notification configurations.
	SlackConfig []*SlackConfig `protobuf:"bytes,6,rep,name=slack_config" json:"slack_config,omitempty"`
	// Zero or more Flowdock notification configurations.
	FlowdockConfig   []*FlowdockConfig `protobuf:"bytes,7,rep,name=flowdock_config" json:"flowdock_config,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *NotificationConfig) Reset()         { *m = NotificationConfig{} }
func (m *NotificationConfig) String() string { return proto.CompactTextString(m) }
func (*NotificationConfig) ProtoMessage()    {}

func (m *NotificationConfig) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *NotificationConfig) GetPagerdutyConfig() []*PagerDutyConfig {
	if m != nil {
		return m.PagerdutyConfig
	}
	return nil
}

func (m *NotificationConfig) GetEmailConfig() []*EmailConfig {
	if m != nil {
		return m.EmailConfig
	}
	return nil
}

func (m *NotificationConfig) GetPushoverConfig() []*PushoverConfig {
	if m != nil {
		return m.PushoverConfig
	}
	return nil
}

func (m *NotificationConfig) GetHipchatConfig() []*HipChatConfig {
	if m != nil {
		return m.HipchatConfig
	}
	return nil
}

func (m *NotificationConfig) GetSlackConfig() []*SlackConfig {
	if m != nil {
		return m.SlackConfig
	}
	return nil
}

func (m *NotificationConfig) GetFlowdockConfig() []*FlowdockConfig {
	if m != nil {
		return m.FlowdockConfig
	}
	return nil
}

// A regex-based label filter used in aggregations.
type Filter struct {
	// The regex matching the label name.
	NameRe *string `protobuf:"bytes,1,opt,name=name_re" json:"name_re,omitempty"`
	// The regex matching the label value.
	ValueRe          *string `protobuf:"bytes,2,opt,name=value_re" json:"value_re,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}

func (m *Filter) GetNameRe() string {
	if m != nil && m.NameRe != nil {
		return *m.NameRe
	}
	return ""
}

func (m *Filter) GetValueRe() string {
	if m != nil && m.ValueRe != nil {
		return *m.ValueRe
	}
	return ""
}

// Grouping and notification setting definitions for alerts.
type AggregationRule struct {
	// Filters that define which alerts are matched by this AggregationRule.
	Filter []*Filter `protobuf:"bytes,1,rep,name=filter" json:"filter,omitempty"`
	// How many seconds to wait before resending a notification for a specific alert.
	RepeatRateSeconds *int32 `protobuf:"varint,2,opt,name=repeat_rate_seconds,def=7200" json:"repeat_rate_seconds,omitempty"`
	// Notification configuration to use for this AggregationRule, referenced by
	// their name.
	NotificationConfigName *string `protobuf:"bytes,3,opt,name=notification_config_name" json:"notification_config_name,omitempty"`
	XXX_unrecognized       []byte  `json:"-"`
}

func (m *AggregationRule) Reset()         { *m = AggregationRule{} }
func (m *AggregationRule) String() string { return proto.CompactTextString(m) }
func (*AggregationRule) ProtoMessage()    {}

const Default_AggregationRule_RepeatRateSeconds int32 = 7200

func (m *AggregationRule) GetFilter() []*Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *AggregationRule) GetRepeatRateSeconds() int32 {
	if m != nil && m.RepeatRateSeconds != nil {
		return *m.RepeatRateSeconds
	}
	return Default_AggregationRule_RepeatRateSeconds
}

func (m *AggregationRule) GetNotificationConfigName() string {
	if m != nil && m.NotificationConfigName != nil {
		return *m.NotificationConfigName
	}
	return ""
}

// An InhibitRule specifies that a class of (source) alerts should inhibit
// notifications for another class of (target) alerts if all specified matching
// labels are equal between the two alerts. This may be used to inhibit alerts
// from sending notifications if their meaning is logically a subset of a
// higher-level alert.
//
// For example, if an entire job is down, there is little sense in sending a
// notification for every single instance of said job being down. This could be
// expressed as the following inhibit rule:
//
// inhibit_rule {
//   # Select all source alerts that are candidates for being inhibitors. All
//   # supplied source filters have to match in order to select a source alert.
//   source_filter: {
//     name_re: "alertname"
//     value_re: "JobDown"
//   }
//   source_filter: {
//     name_re: "service"
//     value_re: "api"
//   }
//
//   # Select all target alerts that are candidates for being inhibited. All
//   # supplied target filters have to match in order to select a target alert.
//   target_filter: {
//     name_re: "alertname"
//     value_re: "InstanceDown"
//   }
//   target_filter: {
//     name_re: "service"
//     value_re: "api"
//   }
//
//   # A target alert only actually inhibits a source alert if they match on
//   # these labels. I.e. the alerts needs to fire for the same job in the same
//   # zone for the inhibit to take effect between them.
//   match_on: "job"
//   match_on: "zone"
// }
//
// In this example, when JobDown is firing for
//
//   JobDown{zone="aa",job="test",service="api"}
//
// ...it would inhibit an InstanceDown alert for
//
//   InstanceDown{zone="aa",job="test",instance="1",service="api"}
//
// However, an InstanceDown alert for another zone:
//
//   {zone="ab",job="test",instance="1",service="api"}
//
// ...would still fire.
type InhibitRule struct {
	// The set of Filters which define the group of source alerts (which inhibit
	// the target alerts).
	SourceFilter []*Filter `protobuf:"bytes,1,rep,name=source_filter" json:"source_filter,omitempty"`
	// The set of Filters which define the group of target alerts (which are
	// inhibited by the source alerts).
	TargetFilter []*Filter `protobuf:"bytes,2,rep,name=target_filter" json:"target_filter,omitempty"`
	// A set of label names whose label values need to be identical in source and
	// target alerts in order for the inhibition to take effect.
	MatchOn          []string `protobuf:"bytes,3,rep,name=match_on" json:"match_on,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *InhibitRule) Reset()         { *m = InhibitRule{} }
func (m *InhibitRule) String() string { return proto.CompactTextString(m) }
func (*InhibitRule) ProtoMessage()    {}

func (m *InhibitRule) GetSourceFilter() []*Filter {
	if m != nil {
		return m.SourceFilter
	}
	return nil
}

func (m *InhibitRule) GetTargetFilter() []*Filter {
	if m != nil {
		return m.TargetFilter
	}
	return nil
}

func (m *InhibitRule) GetMatchOn() []string {
	if m != nil {
		return m.MatchOn
	}
	return nil
}

// Global alert manager configuration.
type AlertManagerConfig struct {
	// Aggregation rule definitions.
	AggregationRule []*AggregationRule `protobuf:"bytes,1,rep,name=aggregation_rule" json:"aggregation_rule,omitempty"`
	// Notification configuration definitions.
	NotificationConfig []*NotificationConfig `protobuf:"bytes,2,rep,name=notification_config" json:"notification_config,omitempty"`
	// List of alert inhibition rules.
	InhibitRule      []*InhibitRule `protobuf:"bytes,3,rep,name=inhibit_rule" json:"inhibit_rule,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *AlertManagerConfig) Reset()         { *m = AlertManagerConfig{} }
func (m *AlertManagerConfig) String() string { return proto.CompactTextString(m) }
func (*AlertManagerConfig) ProtoMessage()    {}

func (m *AlertManagerConfig) GetAggregationRule() []*AggregationRule {
	if m != nil {
		return m.AggregationRule
	}
	return nil
}

func (m *AlertManagerConfig) GetNotificationConfig() []*NotificationConfig {
	if m != nil {
		return m.NotificationConfig
	}
	return nil
}

func (m *AlertManagerConfig) GetInhibitRule() []*InhibitRule {
	if m != nil {
		return m.InhibitRule
	}
	return nil
}

func init() {
}
