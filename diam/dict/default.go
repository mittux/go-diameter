// Copyright 2013-2015 go-diameter authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// This file is auto-generated from our dictionaries.

package dict

import (
	"bytes"
	"fmt"
)

// Default is a Parser object with pre-loaded
// Base Protocol and Credit Control dictionaries.
var Default *Parser

func init() {
	var dictionaries = []struct{ name, xml string }{
		{"Base", baseXML},
		{"Credit Control", creditcontrolXML},
		{"Gx Charging Control", gxcreditcontrolXML},
		{"Network Access Server", networkaccessserverXML},
		/*	{"TGPP", tgpprorfXML},  */
		{"TGPP_S6a", tgpps6aXML},
		{"TGPP_CLF", tgppe2XML},
	}
	var err error
	Default, err = NewParser()
	if err != nil {
		panic(err)
	}
	for _, dict := range dictionaries {
		err = Default.Load(bytes.NewReader([]byte(dict.xml)))
		if err != nil {
			panic(fmt.Sprintf("Cannot load %s dictionary: %s", dict.name, err))
		}
	}
}

var baseXML = `<?xml version="1.0" encoding="UTF-8"?>
<diameter>

	<application id="0" name="Base"> <!-- Diameter Common Messages -->

		<command code="257" short="CE" name="Capabilities-Exchange">
			<request>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Host-IP-Address" required="true" min="1"/>
				<rule avp="Vendor-Id" required="true" max="1"/>
				<rule avp="Product-Name" required="true" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Supported-Vendor-Id" required="False"/>
				<rule avp="Auth-Application-Id" required="False"/>
				<rule avp="Inband-Security-Id" required="False"/>
				<rule avp="Acct-Application-Id" required="False"/>
				<rule avp="Vendor-Specific-Application-Id" required="False"/>
				<rule avp="Firmware-Revision" required="False" max="1"/>
			</request>
			<answer>
				<rule avp="Result-Code" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Host-IP-Address" required="true" min="1"/>
				<rule avp="Vendor-Id" required="true" max="1"/>
				<rule avp="Product-Name" required="true" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Error-Message" required="false" max="1"/>
				<rule avp="Failed-AVP" required="false" max="1"/>
				<rule avp="Supported-Vendor-Id" required="False"/>
				<rule avp="Auth-Application-Id" required="False"/>
				<rule avp="Inband-Security-Id" required="False"/>
				<rule avp="Acct-Application-Id" required="False"/>
				<rule avp="Vendor-Specific-Application-Id" required="False"/>
				<rule avp="Firmware-Revision" required="False" max="1"/>
			</answer>
		</command>

		<command code="258" short="RA" name="Re-Auth">
			<request>
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Destination-Realm" required="true" max="1"/>
				<rule avp="Destination-Host" required="true" max="1"/>
				<rule avp="Auth-Application-Id" required="true" max="1"/>
				<rule avp="Re-Auth-Request-Type" required="true" max="1"/>
				<rule avp="User-Name" required="false" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Proxy-Info" required="false"/>
				<rule avp="Route-Record" required="false"/>
			</request>
			<answer>
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Result-Code" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="User-Name" required="false" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Error-Message" required="false" max="1"/>
				<rule avp="Error-Reporting-Host" required="false" max="1"/>
				<rule avp="Failed-AVP" required="false" max="1"/>
				<rule avp="Redirect-Host" required="false"/>
				<rule avp="Redirect-Host-Usage" required="false" max="1"/>
				<rule avp="Redirect-Max-Cache-Time" required="false" max="1"/>
				<rule avp="Proxy-Info" required="false"/>
			</answer>
		</command>

		<command code="271" short="AC" name="Accounting">
			<request>
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Destination-Realm" required="true" max="1"/>
				<rule avp="Accounting-Record-Type" required="true" max="1"/>
				<rule avp="Accounting-Record-Number" required="true" max="1"/>
				<rule avp="Acct-Application-Id" required="false" max="1"/>
				<rule avp="Vendor-Specific-Application-Id" required="false" max="1"/>
				<rule avp="User-Name" required="false" max="1"/>
				<rule avp="Destination-Host" required="false" max="1"/>
				<rule avp="Accounting-Sub-Session-Id" required="false" max="1"/>
				<rule avp="Acct-Session-Id" required="false" max="1"/>
				<rule avp="Acct-Multi-Session-Id" required="false" max="1"/>
				<rule avp="Acct-Interim-Interval" required="false" max="1"/>
				<rule avp="Accounting-Realtime-Required" required="false" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Event-Timestamp" required="false" max="1"/>
				<rule avp="Proxy-Info" required="false"/>
				<rule avp="Route-Record" required="false"/>
			</request>
			<answer>
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Result-Code" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Accounting-Record-Type" required="true" max="1"/>
				<rule avp="Accounting-Record-Number" required="true" max="1"/>
				<rule avp="Acct-Application-Id" required="false" max="1"/>
				<rule avp="Vendor-Specific-Application-Id" required="false" max="1"/>
				<rule avp="User-Name" required="false" max="1"/>
				<rule avp="Accounting-Sub-Session-Id" required="false" max="1"/>
				<rule avp="Acct-Session-Id" required="false" max="1"/>
				<rule avp="Acct-Multi-Session-Id" required="false" max="1"/>
				<rule avp="Error-Message" required="false" max="1"/>
				<rule avp="Error-Reporting-Host" required="false" max="1"/>
				<rule avp="Failed-AVP" required="false" max="1"/>
				<rule avp="Acct-Interim-Interval" required="false" max="1"/>
				<rule avp="Accounting-Realtime-Required" required="false" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Event-Timestamp" required="false" max="1"/>
				<rule avp="Proxy-Info" required="false"/>
			</answer>
		</command>

		<command code="274" short="AS" name="Abort-Session">
			<request>
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Destination-Realm" required="true" max="1"/>
				<rule avp="Destination-Host" required="true" max="1"/>
				<rule avp="Auth-Application-Id" required="true" max="1"/>
				<rule avp="User-Name" required="false" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Proxy-Info" required="false"/>
				<rule avp="Route-Record" required="false"/>
			</request>
			<answer>
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Result-Code" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="User-Name" required="false" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Error-Message" required="false" max="1"/>
				<rule avp="Error-Reporting-Host" required="false" max="1"/>
				<rule avp="Failed-AVP" required="false" max="1"/>
				<rule avp="Redirect-Host" required="false"/>
				<rule avp="Redirect-Host-Usage" required="false" max="1"/>
				<rule avp="Redirect-Max-Cache-Time" required="false" max="1"/>
				<rule avp="Proxy-Info" required="false"/>
			</answer>
		</command>

		<command code="275" short="ST" name="Session-Termination">
			<request>
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Destination-Realm" required="true" max="1"/>
				<rule avp="Auth-Application-Id" required="true" max="1"/>
				<rule avp="Termination-Cause" required="true" max="1"/>
				<rule avp="User-Name" required="false" max="1"/>
				<rule avp="Destination-Host" required="false" max="1"/>
				<rule avp="Class" required="false"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Proxy-Info" required="false"/>
				<rule avp="Route-Record" required="false"/>
			</request>
			<answer>
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Result-Code" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="User-Name" required="false" max="1"/>
				<rule avp="Class" required="false"/>
				<rule avp="Error-Message" required="false" max="1"/>
				<rule avp="Error-Reporting-Host" required="false" max="1"/>
				<rule avp="Failed-AVP" required="false" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Redirect-Host" required="false"/>
				<rule avp="Redirect-Host-Usage" required="false" max="1"/>
				<rule avp="Redirect-Max-Cache-Time" required="false" max="1"/>
				<rule avp="Proxy-Info" required="false"/>
			</answer>
		</command>

		<command code="280" short="DW" name="Device-Watchdog">
			<request>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
			</request>
			<answer>
				<rule avp="Result-Code" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Error-Message" required="false" max="1"/>
				<rule avp="Failed-AVP" required="false" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
			</answer>
		</command>

		<command code="282" short="DP" name="Disconnect-Peer">
			<request>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Disconnect-Cause" required="false" max="1"/>
			</request>
			<answer>
				<rule avp="Result-Code" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Error-Message" required="false" max="1"/>
				<rule avp="Failed-AVP" required="false" max="1"/>
			</answer>
		</command>

		<avp name="Acct-Interim-Interval" code="85" must="M" may="P" must-not="V" may-encrypt="Y">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Accounting-Realtime-Required" code="483" must="M" may="P" must-not="V" may-encrypt="Y">
			<data type="Enumerated">
				<item code="1" name="DELIVER_AND_GRANT"/>
				<item code="2" name="GRANT_AND_STORE"/>
				<item code="3" name="GRANT_AND_LOSE"/>
			</data>
		</avp>

		<avp name="Acct-Multi-Session-Id" code="50" must="M" may="P" must-not="V" may-encrypt="Y">
			<data type="UTF8String"/>
		</avp>

		<avp name="Accounting-Record-Number" code="485" must="M" may="P" must-not="V" may-encrypt="Y">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Accounting-Record-Type" code="480" must="M" may="P" must-not="V" may-encrypt="Y">
			<data type="Enumerated">
				<item code="1" name="EVENT_RECORD"/>
				<item code="2" name="START_RECORD"/>
				<item code="3" name="INTERIM_RECORD"/>
				<item code="4" name="STOP_RECORD"/>
			</data>
		</avp>

		<avp name="Accounting-Session-Id" code="44" must="M" may="P" must-not="V" may-encrypt="Y">
			<data type="OctetString"/>
		</avp>

		<avp name="Accounting-Sub-Session-Id" code="287" must="M" may="P" must-not="V" may-encrypt="Y">
			<data type="Unsigned64"/>
		</avp>

		<avp name="Acct-Application-Id" code="259" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Auth-Application-Id" code="258" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Auth-Request-Type" code="274" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Enumerated">
				<item code="1" name="AUTHENTICATE_ONLY"/>
				<item code="2" name="AUTHORIZE_ONLY"/>
				<item code="3" name="AUTHORIZE_AUTHENTICATE"/>
			</data>
		</avp>

		<avp name="Authorization-Lifetime" code="291" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Auth-Grace-Period" code="276" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Auth-Session-State" code="277" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Enumerated">
				<item code="0" name="STATE_MAINTAINED"/>
				<item code="1" name="NO_STATE_MAINTAINED"/>
			</data>
		</avp>

		<avp name="Re-Auth-Request-Type" code="285" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Enumerated">
				<item code="0" name="AUTHORIZE_ONLY"/>
				<item code="1" name="AUTHORIZE_AUTHENTICATE"/>
			</data>
		</avp>

		<avp name="Class" code="25" must="M" may="P" must-not="V" may-encrypt="Y">
			<data type="OctetString"/>
		</avp>

		<avp name="Destination-Host" code="293" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="DiameterIdentity"/>
		</avp>

		<avp name="Destination-Realm" code="283" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="DiameterIdentity"/>
		</avp>

		<avp name="Disconnect-Cause" code="273" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Enumerated">
				<item code="0" name="REBOOTING"/>
				<item code="1" name="BUSY"/>
				<item code="2" name="DO_NOT_WANT_TO_TALK_TO_YOU"/>
			</data>
		</avp>

		<avp name="Error-Message" code="281" must="-" may="P" must-not="V,M" may-encrypt="-">
			<data type="UTF8String"/>
		</avp>

		<avp name="Error-Reporting-Host" code="294" must="-" may="P" must-not="V,M" may-encrypt="-">
			<data type="DiameterIdentity"/>
		</avp>

		<avp name="Event-Timestamp" code="55" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Time"/>
		</avp>

		<avp name="Experimental-Result" code="297" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Grouped">
				<rule avp="Vendor-Id" required="true" max="1"/>
				<rule avp="Experimental-Result-Code" required="true" max="1"/>
			</data>
		</avp>

		<avp name="Experimental-Result-Code" code="298" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Failed-AVP" code="279" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Grouped"/>
		</avp>

		<avp name="Firmware-Revision" code="267" must="-" may="-" must-not="P,V,M" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Host-IP-Address" code="257" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Address"/>
		</avp>

		<avp name="Inband-Security-Id" code="299" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Multi-Round-Time-Out" code="272" must="M" may="P" must-not="V" may-encrypt="Y">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Origin-Host" code="264" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="DiameterIdentity"/>
		</avp>

		<avp name="Origin-Realm" code="296" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="DiameterIdentity"/>
		</avp>

		<avp name="Origin-State-Id" code="278" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Product-Name" code="269" must="-" may="-" must-not="P,V,M" may-encrypt="-">
			<data type="UTF8String"/>
		</avp>

		<avp name="Proxy-Host" code="280" must="M" may="-" must-not="P,V" may-encrypt="-">
			<data type="DiameterIdentity"/>
		</avp>

		<avp name="Proxy-Info" code="284" must="M" may="-" must-not="P,V" may-encrypt="-">
			<data type="Grouped">
				<rule avp="Proxy-Host" required="true" max="1"/>
				<rule avp="Proxy-State" required="true" max="1"/>
			</data>
		</avp>

		<avp name="Proxy-State" code="33" must="M" may="-" must-not="P,V" may-encrypt="-">
			<data type="OctetString"/>
		</avp>

		<avp name="Redirect-Host" code="292" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="DiameterURI"/>
		</avp>

		<avp name="Redirect-Host-Usage" code="261" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Enumerated">
				<item code="0" name="DONT_CACHE"/>
				<item code="1" name="ALL_SESSION"/>
				<item code="2" name="ALL_REALM"/>
				<item code="3" name="REALM_AND_APPLICATION"/>
				<item code="4" name="ALL_APPLICATION"/>
				<item code="5" name="ALL_HOST"/>
				<item code="6" name="ALL_USER"/>
			</data>
		</avp>

		<avp name="Redirect-Max-Cache-Time" code="262" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Result-Code" code="268" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Route-Record" code="282" must="M" may="-" must-not="P,V" may-encrypt="-">
			<data type="DiameterIdentity"/>
		</avp>

		<avp name="Session-Id" code="263" must="M" may="P" must-not="V" may-encrypt="Y">
			<data type="UTF8String"/>
		</avp>

		<avp name="Session-Timeout" code="27" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Session-Binding" code="270" must="M" may="P" must-not="V" may-encrypt="Y">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Session-Server-Failover" code="271" must="M" may="P" must-not="V" may-encrypt="Y">
			<data type="Enumerated">
				<item code="0" name="REFUSE_SERVICE"/>
				<item code="1" name="TRY_AGAIN"/>
				<item code="2" name="ALLOW_SERVICE"/>
				<item code="3" name="TRY_AGAIN_ALLOW_SERVICE"/>
			</data>
		</avp>

		<avp name="Supported-Vendor-Id" code="265" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Termination-Cause" code="295" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Enumerated">
				<item code="1" name="DIAMETER_LOGOUT"/>
				<item code="2" name="DIAMETER_SERVICE_NOT_PROVIDED"/>
				<item code="3" name="DIAMETER_BAD_ANSWER"/>
				<item code="4" name="DIAMETER_ADMINISTRATIVE"/>
				<item code="5" name="DIAMETER_LINK_BROKEN"/>
				<item code="6" name="DIAMETER_AUTH_EXPIRED"/>
				<item code="7" name="DIAMETER_USER_MOVED"/>
				<item code="8" name="DIAMETER_SESSION_TIMEOUT"/>
			</data>
		</avp>

		<avp name="User-Name" code="1" must="M" may="P" must-not="V" may-encrypt="Y">
			<data type="UTF8String"/>
		</avp>

		<avp name="Vendor-Id" code="266" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Vendor-Specific-Application-Id" code="260" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Grouped">
				<rule avp="Vendor-Id" required="false" max="1"/>
				<rule avp="Auth-Application-Id" required="true" max="1"/>
				<rule avp="Acct-Application-Id" required="true" max="1"/>
			</data>
		</avp>

		<!-- IETF RFC 7683 - https://tools.ietf.org/html/rfc7683 -->
		<avp name="OC-Supported-Features" code="621" must-not="V">
			<data type="Grouped">
				<rule avp="OC-Feature-Vector" required="false"/>
				<rule avp="AVP" required="false"/>
			</data>
		</avp>

		<avp name="OC-Feature-Vector" code="622" must-not="V">
			<data type="Unsigned64"/>
		</avp>

		<avp name="OC-OLR" code="623" must-not="V">
			<data type="Grouped">
				<rule avp="OC-Sequence-Number" required="true" max="1"/>
				<rule avp="OC-Report-Type" required="true" max="1"/>
				<rule avp="OC-Reduction-Percentage" required="false" max="1"/>
				<rule avp="OC-Validity-Duration" required="false" max="1"/>
				<rule avp="AVP" required="false"/>
			</data>
		</avp>

		<avp name="OC-Sequence-Number" code="624" must-not="V">
			<data type="Unsigned64"/>
		</avp>

		<avp name="OC-Validity-Duration" code="625" must-not="V">
			<data type="Unsigned32"/>
		</avp>

		<avp name="OC-Report-Type" code="626" must-not="V">
			<data type="Enumerated">
				<item code="0" name="HOST_REPORT"/>
				<item code="1" name="REALM_REPORT"/>
			</data>
		</avp>

		<avp name="OC-Reduction-Percentage" code="627" must-not="V">
			<data type="Unsigned32"/>
		</avp>

		<!-- IETF RFC 7944 - https://tools.ietf.org/html/rfc7944 -->
		<avp name="DRMP" code="301" must-not="V">
			<data type="Enumerated">
				<item code="0" name="PRIORITY_0"/>
				<item code="1" name="PRIORITY_1"/>
				<item code="2" name="PRIORITY_2"/>
				<item code="3" name="PRIORITY_3"/>
				<item code="4" name="PRIORITY_4"/>
				<item code="5" name="PRIORITY_5"/>
				<item code="6" name="PRIORITY_6"/>
				<item code="7" name="PRIORITY_7"/>
				<item code="8" name="PRIORITY_8"/>
				<item code="9" name="PRIORITY_9"/>
				<item code="10" name="PRIORITY_10"/>
				<item code="11" name="PRIORITY_11"/>
				<item code="12" name="PRIORITY_12"/>
				<item code="13" name="PRIORITY_13"/>
				<item code="14" name="PRIORITY_14"/>
				<item code="15" name="PRIORITY_15"/>
			</data>
		</avp>


	</application>
	<application id="3" type="acct" name="Base Accounting"> <!-- Diameter Base Accounting Messages -->
	</application>
</diameter>`

var creditcontrolXML = `<?xml version="1.0" encoding="UTF-8"?>
<diameter>

	<application id="4" type="auth" name="Charging Control">
		<!-- Diameter Credit Control Application -->
		<!-- http://tools.ietf.org/html/rfc4006 -->

		<command code="272" short="CC" name="Credit-Control">
			<request>
				<!-- http://tools.ietf.org/html/rfc4006#section-3.1 -->
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Destination-Realm" required="true" max="1"/>
				<rule avp="Auth-Application-Id" required="true" max="1"/>
				<rule avp="Service-Context-Id" required="true" max="1"/>
				<rule avp="CC-Request-Type" required="true" max="1"/>
				<rule avp="CC-Request-Number" required="true" max="1"/>
				<rule avp="Destination-Host" required="false" max="1"/>
				<rule avp="User-Name" required="false" max="1"/>
				<rule avp="CC-Sub-Session-Id" required="false" max="1"/>
				<rule avp="Acct-Multi-Session-Id" required="false" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Event-Timestamp" required="false" max="1"/>
				<rule avp="Subscription-Id" required="false" max="1"/>
				<rule avp="Service-Identifier" required="false" max="1"/>
				<rule avp="Termination-Cause" required="false" max="1"/>
				<rule avp="Requested-Service-Unit" required="false" max="1"/>
				<rule avp="Requested-Action" required="false" max="1"/>
				<rule avp="Used-Service-Unit" required="false" max="1"/>
				<rule avp="Multiple-Services-Indicator" required="false" max="1"/>
				<rule avp="Multiple-Services-Credit-Control" required="false" max="1"/>
				<rule avp="Service-Parameter-Info" required="false" max="1"/>
				<rule avp="CC-Correlation-Id" required="false" max="1"/>
				<rule avp="User-Equipment-Info" required="false" max="1"/>
				<rule avp="Proxy-Info" required="false" max="1"/>
				<rule avp="Route-Record" required="false" max="1"/>
			</request>
			<answer>
				<!-- http://tools.ietf.org/html/rfc4006#section-3.2 -->
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Result-Code" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="CC-Request-Type" required="true" max="1"/>
				<rule avp="CC-Request-Number" required="true" max="1"/>
				<rule avp="User-Name" required="false" max="1"/>
				<rule avp="CC-Session-Failover" required="false" max="1"/>
				<rule avp="CC-Sub-Session-Id" required="false" max="1"/>
				<rule avp="Acct-Multi-Session-Id" required="false" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Event-Timestamp" required="false" max="1"/>
				<rule avp="Granted-Service-Unit" required="false" max="1"/>
				<rule avp="Multiple-Services-Credit-Control" required="false" max="1"/>
				<rule avp="Cost-Information" required="false" max="1"/>
				<rule avp="Final-Unit-Indication" required="false" max="1"/>
				<rule avp="Check-Balance-Result" required="false" max="1"/>
				<rule avp="Credit-Control-Failure-Handling" required="false" max="1"/>
				<rule avp="Direct-Debiting-Failure-Handling" required="false" max="1"/>
				<rule avp="Validity-Time" required="false" max="1"/>
				<rule avp="Redirect-Host" required="false" max="1"/>
				<rule avp="Redirect-Host-Usage" required="false" max="1"/>
				<rule avp="Redirect-Max-Cache-Time" required="false" max="1"/>
				<rule avp="Proxy-Info" required="false" max="1"/>
				<rule avp="Route-Record" required="false" max="1"/>
				<rule avp="Failed-AVP" required="false" max="1"/>
			</answer>
		</command>

		<avp name="CC-Correlation-Id" code="411" must="-" may="P,M" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.1 -->
			<data type="OctetString"/>
		</avp>

		<avp name="CC-Input-Octets" code="412" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.24 -->
			<data type="Unsigned64"/>
		</avp>

		<avp name="CC-Money" code="413" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.22 -->
			<data type="Grouped">
				<rule avp="Unit-Value" required="true" max="1"/>
				<rule avp="Currency-Code" required="true" max="1"/>
			</data>
		</avp>

		<avp name="CC-Output-Octets" code="414" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.25 -->
			<data type="Unsigned64"/>
		</avp>

		<avp name="CC-Request-Number" code="415" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.2 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="CC-Request-Type" code="416" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.3 -->
			<data type="Enumerated">
				<item code="1" name="INITIAL_REQUEST"/>
				<item code="2" name="UPDATE_REQUEST"/>
				<item code="3" name="TERMINATION_REQUEST"/>
			</data>
		</avp>

		<avp name="CC-Service-Specific-Units" code="417" must="M" may="P" must-not="V" may-encrypt="Y">
			<data type="Unsigned64"/>
		</avp>

		<avp name="CC-Session-Failover" code="418" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.4 -->
			<data type="Enumerated">
				<item code="0" name="FAILOVER_NOT_SUPPORTED"/>
				<item code="1" name="FAILOVER_SUPPORTED"/>
			</data>
		</avp>

		<avp name="CC-Sub-Session-Id" code="419" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.5 -->
			<data type="Unsigned64"/>
		</avp>

		<avp name="CC-Time" code="420" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.21 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="CC-Total-Octets" code="421" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.23 -->
			<data type="Unsigned64"/>
		</avp>

		<avp name="CC-Unit-Type" code="454" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.32 -->
			<data type="Enumerated">
				<item code="0" name="TIME"/>
				<item code="1" name="MONEY"/>
				<item code="2" name="TOTAL-OCTETS"/>
				<item code="3" name="INPUT-OCTETS"/>
				<item code="4" name="OUTPUT-OCTETS"/>
				<item code="5" name="SERVICE-SPECIFIC-UNITS"/>
			</data>
		</avp>

		<avp name="Check-Balance-Result" code="422" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.6 -->
			<data type="Enumerated">
				<item code="0" name="ENOUGH_CREDIT"/>
				<item code="1" name="NO_CREDIT"/>
			</data>
		</avp>

		<avp name="Cost-Information" code="423" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.7 -->
			<data type="Grouped">
				<rule avp="Unit-Value" required="true" max="1"/>
				<rule avp="Currency-Code" required="true" max="1"/>
				<rule avp="Cost-Unit" required="true" max="1"/>
			</data>
		</avp>

		<avp name="Cost-Unit" code="424" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.12 -->
			<data type="UTF8String"/>
		</avp>

		<avp name="Credit-Control" code="426" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.13 -->
			<data type="Enumerated">
				<item code="0" name="CREDIT_AUTHORIZATION"/>
				<item code="1" name="RE_AUTHORIZATION"/>
			</data>
		</avp>

		<avp name="Credit-Control-Failure-Handling" code="427" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.14 -->
			<data type="Enumerated">
				<item code="0" name="TERMINATE"/>
				<item code="1" name="CONTINUE"/>
				<item code="2" name="RETRY_AND_TERMINATE"/>
			</data>
		</avp>

		<avp name="Currency-Code" code="425" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.11 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="Direct-Debiting-Failure-Handling" code="428" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.15 -->
			<data type="Enumerated">
				<item code="0" name="TERMINATE_OR_BUFFER"/>
				<item code="1" name="CONTINUE"/>
			</data>
		</avp>

		<avp name="Exponent" code="429" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.9 -->
			<data type="Integer32"/>
		</avp>

		<avp name="Final-Unit-Action" code="449" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.35 -->
			<data type="Enumerated">
				<item code="0" name="TERMINATE"/>
				<item code="1" name="REDIRECT"/>
				<item code="2" name="RESTRICT_ACCESS"/>
			</data>
		</avp>

		<avp name="Final-Unit-Indication" code="430" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.34 -->
			<data type="Grouped">
				<rule avp="Final-Unit-Action" required="true" max="1"/>
				<rule avp="Restriction-Filter-Rule" required="false" max="1"/>
				<rule avp="Filter-Id" required="false" max="1"/>
				<rule avp="Redirect-Server" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Granted-Service-Unit" code="431" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.17 -->
			<data type="Grouped">
				<rule avp="Tariff-Time-Change" required="false" max="1"/>
				<rule avp="CC-Time" required="false" max="1"/>
				<rule avp="CC-Money" required="false" max="1"/>
				<rule avp="CC-Total-Octets" required="false" max="1"/>
				<rule avp="CC-Input-Octets" required="false" max="1"/>
				<rule avp="CC-Output-Octets" required="false" max="1"/>
				<rule avp="CC-Service-Specific-Units" required="false" max="1"/>
				<!-- *[ AVP ]-->
			</data>
		</avp>

		<avp name="G-S-U-Pool-Identifier" code="453" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.31 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="G-S-U-Pool-Reference" code="457" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.30 -->
			<data type="Grouped">
				<rule avp="G-S-U-Pool-Identifier" required="true" max="1"/>
				<rule avp="CC-Unit-Type" required="true" max="1"/>
				<rule avp="Unit-Value" required="true" max="1"/>
			</data>
		</avp>

		<avp name="Multiple-Services-Credit-Control" code="456" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.16 -->
			<data type="Grouped">
				<rule avp="Granted-Service-Unit" required="false" max="1"/>
				<rule avp="Requested-Service-Unit" required="false" max="1"/>
				<rule avp="Used-Service-Unit" required="false" max="1"/>
				<rule avp="Tariff-Change-Usage" required="false" max="1"/>
				<rule avp="Service-Identifier" required="false" max="1"/>
				<rule avp="Rating-Group" required="false" max="1"/>
				<rule avp="G-S-U-Pool-Reference" required="false" max="1"/>
				<rule avp="Validity-Time" required="false" max="1"/>
				<rule avp="Result-Code" required="false" max="1"/>
				<rule avp="Final-Unit-Indication" required="false" max="1"/>
				<!-- *[ AVP ]-->
			</data>
		</avp>

		<avp name="Multiple-Services-Indicator" code="455" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.40 -->
			<data type="Enumerated">
				<item code="0" name="MULTIPLE_SERVICES_NOT_SUPPORTED"/>
				<item code="1" name="MULTIPLE_SERVICES_SUPPORTED"/>
			</data>
		</avp>

		<avp name="Rating-Group" code="432" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.29 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="Redirect-Address-Type" code="433" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.38 -->
			<data type="Enumerated">
				<item code="0" name="IPv4 Address"/>
				<item code="1" name="IPv6 Address"/>
				<item code="2" name="URL"/>
				<item code="3" name="SIP URI"/>
			</data>
		</avp>

		<avp name="Redirect-Server" code="434" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.37 -->
			<data type="Grouped">
				<rule avp="Redirect-Address-Type" required="true" max="1"/>
				<rule avp="Redirect-Server-Address" required="true" max="1"/>
			</data>
		</avp>

		<avp name="Redirect-Server-Address" code="435" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.39 -->
			<data type="UTF8String"/>
		</avp>

		<avp name="Requested-Action" code="436" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.41 -->
			<data type="Enumerated">
				<item code="0" name="DIRECT_DEBITING"/>
				<item code="1" name="REFUND_ACCOUNT"/>
				<item code="2" name="CHECK_BALANCE"/>
				<item code="3" name="PRICE_ENQUIRY"/>
			</data>
		</avp>

		<avp name="Requested-Service-Unit" code="437" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.18-->
			<data type="Grouped">
				<rule avp="CC-Time" required="false" max="1"/>
				<rule avp="CC-Money" required="false" max="1"/>
				<rule avp="CC-Total-Octets" required="false" max="1"/>
				<rule avp="CC-Input-Octets" required="false" max="1"/>
				<rule avp="CC-Output-Octets" required="false" max="1"/>
				<rule avp="CC-Service-Specific-Units" required="false" max="1"/>
				<!-- *[ AVP ]-->
			</data>
		</avp>

		<avp name="Restriction-Filter-Rule" code="438" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.36-->
			<data type="IPFilterRule"/>
		</avp>

		<avp name="Service-Context-Id" code="461" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.42-->
			<data type="UTF8String"/>
		</avp>

		<avp name="Service-Identifier" code="439" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.28-->
			<data type="Unsigned32"/>
		</avp>

		<avp name="Service-Parameter-Info" code="440" must="-" may="P,M" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.43-->
			<data type="Grouped">
				<rule avp="Service-Parameter-Type" required="true" max="1"/>
				<rule avp="Service-Parameter-Value" required="true" max="1"/>
			</data>
		</avp>

		<avp name="Service-Parameter-Type" code="441" must="-" may="P,M" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.44-->
			<data type="Unsigned32"/>
		</avp>

		<avp name="Service-Parameter-Value" code="442" must="-" may="P,M" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.45-->
			<data type="OctetString"/>
		</avp>

		<avp name="Subscription-Id" code="443" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.46-->
			<data type="Grouped">
				<rule avp="Subscription-Id-Type" required="true" max="1"/>
				<rule avp="Subscription-Id-Data" required="true" max="1"/>
			</data>
		</avp>

		<avp name="Subscription-Id-Data" code="444" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.48-->
			<data type="UTF8String"/>
		</avp>

		<avp name="Subscription-Id-Type" code="450" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.47-->
			<data type="Enumerated">
				<item code="0" name="END_USER_E164"/>
				<item code="1" name="END_USER_IMSI"/>
				<item code="2" name="END_USER_SIP_URI"/>
				<item code="3" name="END_USER_NAI"/>
			</data>
		</avp>

		<avp name="Tariff-Change-Usage" code="452" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.27-->
			<data type="Enumerated">
				<item code="0" name="UNIT_BEFORE_TARIFF_CHANGE"/>
				<item code="1" name="UNIT_AFTER_TARIFF_CHANGE"/>
				<item code="2" name="UNIT_INDETERMINATE"/>
			</data>
		</avp>

		<avp name="Tariff-Time-Change" code="451" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.20-->
			<data type="Time"/>
		</avp>

		<avp name="Unit-Value" code="445" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.8-->
			<data type="Grouped">
				<rule avp="Value-Digits" required="true" max="1"/>
				<rule avp="Exponent" required="true" max="1"/>
			</data>
		</avp>

		<avp name="Used-Service-Unit" code="446" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.19-->
			<data type="Grouped">
				<rule avp="Tariff-Change-Usage" required="false" max="1"/>
				<rule avp="CC-Time" required="false" max="1"/>
				<rule avp="CC-Money" required="false" max="1"/>
				<rule avp="CC-Total-Octets" required="false" max="1"/>
				<rule avp="CC-Input-Octets" required="false" max="1"/>
				<rule avp="CC-Output-Octets" required="false" max="1"/>
				<rule avp="CC-Service-Specific-Units" required="false" max="1"/>
				<!-- *[ AVP ]-->
			</data>
		</avp>

		<avp name="User-Equipment-Info" code="458" must="-" may="P,M" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.49-->
			<data type="Grouped">
				<rule avp="User-Equipment-Info-Type" required="true" max="1"/>
				<rule avp="User-Equipment-Info-Value" required="true" max="1"/>
			</data>
		</avp>

		<avp name="User-Equipment-Info-Type" code="459" must="-" may="P,M" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.50-->
			<data type="Enumerated">
				<item code="0" name="IMEISV"/>
				<item code="1" name="MAC"/>
				<item code="2" name="EUI64"/>
				<item code="3" name="MODIFIED_EUI64"/>
			</data>
		</avp>

		<avp name="User-Equipment-Info-Value" code="460" must="-" may="P,M" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.51-->
			<data type="OctetString"/>
		</avp>

		<avp name="Value-Digits" code="447" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.10-->
			<data type="Integer64"/>
		</avp>

		<avp name="Validity-Time" code="448" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.33-->
			<data type="Unsigned32"/>
		</avp>
	</application>
</diameter>`

var gxcreditcontrolXML = `<?xml version="1.0" encoding="UTF-8"?>
<diameter>

    <application id="16777238" type="auth" name="Gx Charging Control">
        <!-- Diameter Gx Credit Control Application -->
        <!-- 3GPP 29.212 -->

        <vendor id="10415" name="TGPP"/>
        <command code="272" short="CC" name="Credit-Control">
            <request>
                <!-- 3GPP 29.212 Section 5.6.2 -->
                <rule avp="Session-Id" required="true" max="1"/>
                <rule avp="Origin-Host" required="true" max="1"/>
                <rule avp="Origin-Realm" required="true" max="1"/>
                <rule avp="Destination-Realm" required="true" max="1"/>
                <rule avp="Auth-Application-Id" required="true" max="1"/>
                <rule avp="CC-Request-Type" required="true" max="1"/>
                <rule avp="CC-Request-Number" required="true" max="1"/>
                <rule avp="Destination-Host" required="false" max="1"/>
                <rule avp="Origin-State-Id" required="false" max="1"/>
                <rule avp="Subscription-Id" required="false" max="1"/>
                <rule avp="Termination-Cause" required="false" max="1"/>
                <rule avp="User-Equipment-Info" required="false" max="1"/>
                <rule avp="Proxy-Info" required="false" max="1"/>
                <rule avp="Route-Record" required="false" max="1"/>
                <rule avp="Framed-IP-Address" required="false" max="1"/>
                <rule avp="IP-CAN-Type" required="false" max="1"/>
                <rule avp="Called-Station-Id" required="false" max="1"/>
                <rule avp="RAT-Type" required="false" max="1"/>
            </request>
            <answer>
                <!-- 3GPP 29.212 Section 5.6.3 -->
                <rule avp="Session-Id" required="true" max="1"/>
                <rule avp="Result-Code" required="true" max="1"/>
                <rule avp="Origin-Host" required="true" max="1"/>
                <rule avp="Origin-Realm" required="true" max="1"/>
                <rule avp="CC-Request-Type" required="true" max="1"/>
                <rule avp="CC-Request-Number" required="true" max="1"/>
                <rule avp="Origin-State-Id" required="false" max="1"/>
                <rule avp="Proxy-Info" required="false" max="1"/>
                <rule avp="Route-Record" required="false" max="1"/>
                <rule avp="Failed-AVP" required="false" max="1"/>
                <rule avp="Charging-Rule-Install" required="false"/>
                <rule avp="Usage-Monitoring-Information" required="false"/>
                <rule avp="Event-Trigger" required="false"/>
            </answer>
        </command>

        <avp name="Flow-Description" code="507" must="M,V" may="P" may-encrypt="y" vendor-id="10415">
            <!-- 3GPP 29.212 -->
            <data type="IPFilterRule"/>
        </avp>


        <avp name="Charging-Rule-Install" code="1001" must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
            <!-- 3GPP 29.212 Section 5.3.2 -->
            <data type="Grouped">
                <rule avp="Charging-Rule-Name" required="false"/>
                <rule avp="Charging-Rule-Base-Name" required="false"/>
                <rule avp="Charging-Rule-Definition" required="false"/>
                <rule avp="Rule-Activation-Time" required="false"/>
                <rule avp="Rule-Deactivation-Time" required="false"/>
                <!-- *[ AVP ]-->
            </data>
        </avp>

        <avp name="Charging-Rule-Definition" code="1003" must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
            <!-- 3GPP 29.212 -->
            <data type="Grouped">
                <rule avp="Charging-Rule-Name" required="true" max="1"/>
                <rule avp="Rating-Group" required="false" max="1"/>
                <rule avp="Flow-Description" required="false"/>
                <rule avp="Precedence" required="false" max="1"/>
                <rule avp="Monitoring-Key" required="false" max="1"/>
                <rule avp="Redirect-Information" required="false" max="1"/>
                <!-- *[ AVP ]-->
            </data>
        </avp>

        <avp name="Charging-Rule-Base-Name" code="1004" must="M,V" may="P" may-encrypt="y" vendor-id="10415">
            <!-- 3GPP 29.212 Section 5.3.6 -->
            <data type="UTF8String"/>
        </avp>

        <avp name="Charging-Rule-Name" code="1005" must="M,V" may="P" may-encrypt="y" vendor-id="10415">
            <!-- 3GPP 29.212 Section 5.3.6 -->
            <data type="OctetString"/>
        </avp>

        <avp name="Event-Trigger" code="1006" must="M,V" map="P" may-encrypt="Y" vendor-id="10415">
            <!-- 3GPP 29.212 Section 5.3.7 -->
            <data type="Enumerated">
                <item code="0" name="SGSN_CHANGE"/>
                <item code="1" name="QOS_CHANGE"/>
                <item code="2" name="RAT_CHANGE"/>
                <item code="3" name="TFT_CHANGE"/>
                <item code="4" name="PLMN_CHANGE"/>
                <item code="5" name="LOSS_OF_BEARER"/>
                <item code="6" name="RECOVERY_OF_BEARER"/>
                <item code="7" name="IP-CAN_CHANGE"/>
                <item code="11" name="QOS_CHANGE_EXCEEDING_AUTHORIZATION"/>
                <item code="12" name="RAI_CHANGE"/>
                <item code="13" name="USER_LOCATION_CHANGE"/>
                <item code="14" name="NO_EVENT_TRIGGERS"/>
                <item code="15" name="OUT_OF_CREDIT"/>
                <item code="16" name="REALLOCATION_OF_CREDIT"/>
                <item code="17" name="REVALIDATION_TIMEOUT"/>
                <item code="18" name="UE_IP_ADDRESS_ALLOCATE"/>
                <item code="19" name="UE_IP_ADDRESS_RELEASE"/>
                <item code="20" name="DEFAULT_EPS_BEARER_QOS_CHANGE"/>
                <item code="21" name="AN_GW_CHANGE"/>
                <item code="22" name="SUCCESSFUL_RESOURCE_ALLOCATION"/>
                <item code="23" name="RESOURCE_MODIFICATION_REQUEST"/>
                <item code="24" name="PGW_TRACE_CONTROL"/>
                <item code="25" name="UE_TIME_ZONE_CHANGE"/>
                <item code="26" name="USAGE_REPORT"/>
                <item code="27" name="TAI_CHANGE"/>
                <item code="28" name="ECGI_CHANGE"/>
                <item code="29" name="CHARGING_CORRELATION_EXCHANGE"/>
                <item code="30" name="USER_CSG_INFORMATION_CHANGE"/>
            </data>
        </avp>

        <avp name="Precedence" code="1010" must="M,V" may="P" may-encrypt="y" vendor-id="10415">
            <!-- 3GPP 29.212 -->
            <data type="Unsigned32"/>
        </avp>

        <avp name="IP-CAN-Type" code="1027" must="M,V" map="P" may-encrypt="Y" vendor-id="10415">
            <!-- 3GPP 29.212 Section 5.3.27 -->
            <data type="Enumerated">
                <item code="0" name="3GPP-GPRS"/>
                <item code="1" name="DOCSIS"/>
                <item code="2" name="xDSL"/>
                <item code="3" name="WiMAX"/>
                <item code="4" name="3GPP2"/>
                <item code="5" name="3GPP-EPS"/>
                <item code="6" name="Non-3GPP-EPS"/>
                <item code="7" name="FBA"/>
                <item code="8" name="3GPP-5GS"/>
                <item code="9" name="Non-3GPP-5GS"/>
            </data>
        </avp>

        <avp name="Rule-Activation-Time" code="1043" must="M,V" may="P" may-encrypt="y" vendor-id="10415">
            <!-- 3GPP 29.212 -->
            <data type="Time"/>
        </avp>

        <avp name="Rule-Deactivation-Time" code="1044" must="M,V" may="P" may-encrypt="y" vendor-id="10415">
            <!-- 3GPP 29.212 -->
            <data type="Time"/>
        </avp>

        <avp name="Monitoring-Key" code="1066" must="V" may="P" must-not="M" may-encrypt="y" vendor-id="10415">
            <!-- 3GPP 29.212 -->
            <data type="OctetString"/>
        </avp>

        <avp name="Usage-Monitoring-Information" code="1067" must="V" may="P" must-not="M,V" may-encrypt="y" vendor-id="10415">
            <!-- 3GPP 29.212 -->
            <data type="Grouped">
              <rule avp="Monitoring-Key" required="false" max="1"/>
              <rule avp="Granted-Service-Unit" required="false" max="2"/>
              <rule avp="Used-Service-Unit" required="false" max="2"/>
              <rule avp="Usage-Monitoring-Level" required="false" max="1"/>
            </data>
        </avp>

        <avp name="Usage-Monitoring-Level" code="1068" must="V" may="P" must-not="M" may-encrypt="y" vendor-id="10415">
            <!-- 3GPP 29.212 -->
            <data type="Enumerated">
                <item code="0" name="SESSION_LEVEL"/>
                <item code="1" name="PCC_RULE_LEVEL"/>
            </data>
        </avp>

        <avp name="Redirect-Information" code="1085" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
            <!-- 3GPP 29.212 Section 5.3.82 -->
            <data type="Grouped">
                <rule avp="Redirect-Support" required="true" max="1"/>
                <rule avp="Redirect-Address-Type" required="false" max="1"/>
                <rule avp="Redirect-Server-Address" required="false" max="1"/>
                <!-- *[ AVP ]-->
            </data>
        </avp>

        <avp name="Redirect-Support" code="1086" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
            <!-- 3GPP 29.212 Section 5.3.83 -->
            <data type="Enumerated">
                <item code="0" name="REDIRECTION_DISABLED"/>
                <item code="1" name="REDIRECTION_ENABLED"/>
            </data>
        </avp>

    </application>
</diameter>`

var networkaccessserverXML = `<?xml version="1.0" encoding="UTF-8"?>
<diameter>

	<application id="1" type="auth" name="Network Access">
		<!-- Diameter Network Access Server Application -->
		<!-- http://tools.ietf.org/html/rfc7155 -->

		<command code="265" short="AA" name="AA">
			<request>
				<!-- https://tools.ietf.org/html/rfc7155#section-3.1 -->
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Auth-Application-Id" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Destination-Realm" required="true" max="1"/>
				<rule avp="Auth-Request-Type" required="true" max="1"/>
				<rule avp="Destination-Host" required="false" max="1"/>
				<rule avp="NAS-Identifier" required="false" max="1"/>
				<rule avp="NAS-IP-Address" required="true" max="1"/>
				<rule avp="NAS-IPv6-Address" required="false" max="1"/>
				<rule avp="NAS-Port" required="false" max="1"/>
				<rule avp="NAS-Port-Id" required="false" max="1"/>
				<rule avp="NAS-Port-Type" required="false" max="1"/>
				<rule avp="Origin-AAA-Protocol" required="false" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Port-Limit" required="false" max="1"/>
				<rule avp="User-Name" required="false" max="1"/>
				<rule avp="User-Password" required="false" max="1"/>
				<rule avp="Service-Type" required="false" max="1"/>
				<rule avp="State" required="false" max="1"/>
				<rule avp="Authorization-Lifetime" required="false" max="1"/>
				<rule avp="Auth-Grace-Period" required="false" max="1"/>
				<rule avp="Auth-Session-State" required="false" max="1"/>
				<rule avp="Callback-Number" required="false" max="1"/>
				<rule avp="Called-Station-Id" required="false" max="1"/>
				<rule avp="Calling-Station-Id" required="false" max="1"/>
				<rule avp="Originating-Line-Info" required="false" max="1"/>
				<rule avp="Connect-Info" required="false" max="1"/>
				<rule avp="CHAP-Auth" required="false" max="1"/>
				<rule avp="CHAP-Challenge" required="false" max="1"/>
				<rule avp="Framed-Compression" required="false"/>
				<rule avp="Framed-Interface-Id" required="false" max="1"/>
				<rule avp="Framed-IP-Address" required="false" max="1"/>
				<rule avp="Framed-IPv6-Prefix" required="false"/>
				<rule avp="Framed-IP-Netmask" required="false" max="1"/>
				<rule avp="Framed-MTU" required="false" max="1"/>
				<rule avp="Framed-Protocol" required="false" max="1"/>
				<rule avp="ARAP-Password" required="false" max="1"/>
				<rule avp="ARAP-Security" required="false" max="1"/>
				<rule avp="ARAP-Security-Data" required="false"/>
				<rule avp="Login-IP-Host" required="false"/>
				<rule avp="Login-IPv6-Host" required="false"/>
				<rule avp="Login-LAT-Group" required="false" max="1"/>
				<rule avp="Login-LAT-Node" required="false" max="1"/>
				<rule avp="Login-LAT-Port" required="false" max="1"/>
				<rule avp="Login-LAT-Service" required="false" max="1"/>
				<rule avp="Tunneling" required="false"/>
				<rule avp="Proxy-Info" required="false"/>
				<rule avp="Route-Record" required="false"/>
			</request>
			<answer>
				<!-- http://tools.ietf.org/html/rfc7155#section-3.2 -->
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Auth-Application-Id" required="true" max="1"/>
				<rule avp="Auth-Request-Type" required="true" max="1"/>
				<rule avp="Result-Code" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="User-Name" required="false" max="1"/>
				<rule avp="Service-Type" required="false" max="1"/>
				<rule avp="Class" required="false"/>
				<rule avp="Configuration-Token" required="false"/>
				<rule avp="Acct-Interim-Interval" required="false" max="1"/>
				<rule avp="Error-Message" required="false" max="1"/>
				<rule avp="Error-Reporting-Host" required="false" max="1"/>
				<rule avp="Failed-AVP" required="false"/>
				<rule avp="Idle-Timeout" required="false" max="1"/>
				<rule avp="Authorization-Lifetime" required="false" max="1"/>
				<rule avp="Auth-Grace-Period" required="false" max="1"/>
				<rule avp="Auth-Session-State" required="false" max="1"/>
				<rule avp="Re-Auth-Request-Type" required="false" max="1"/>
				<rule avp="Multi-Round-Time-Out" required="false" max="1"/>
				<rule avp="Session-Timeout" required="false" max="1"/>
				<rule avp="State" required="false" max="1"/>
				<rule avp="Reply-Message" required="false"/>
				<rule avp="Origin-AAA-Protocol" required="false" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Filter-Id" required="false"/>
				<rule avp="Password-Retry" required="false" max="1"/>
				<rule avp="Port-Limit" required="false" max="1"/>
				<rule avp="Prompt" required="false" max="1"/>
				<rule avp="ARAP-Challenge-Response" required="false" max="1"/>
				<rule avp="ARAP-Features" required="false" max="1"/>
				<rule avp="ARAP-Security" required="false" max="1"/>
				<rule avp="ARAP-Security-Data" required="false"/>
				<rule avp="ARAP-Zone-Access" required="false" max="1"/>
				<rule avp="Callback-Id" required="false" max="1"/>
				<rule avp="Callback-Number" required="false" max="1"/>
				<rule avp="Framed-Appletalk-Link" required="false" max="1"/>
				<rule avp="Framed-Appletalk-Network" required="false"/>
				<rule avp="Framed-Appletalk-Zone" required="false" max="1"/>
				<rule avp="Framed-Compression" required="false"/>
				<rule avp="Framed-Interface-Id" required="false" max="1"/>
				<rule avp="Framed-IP-Address" required="false" max="1"/>
				<rule avp="Framed-IPv6-Prefix" required="false"/>
				<rule avp="Framed-IPv6-Pool" required="false" max="1"/>
				<rule avp="Framed-IPv6-Route" required="false"/>
				<rule avp="Framed-IP-Netmask" required="false" max="1"/>
				<rule avp="Framed-Route" required="false"/>
				<rule avp="Framed-Pool" required="false" max="1"/>
				<rule avp="Framed-IPX-Network" required="false" max="1"/>
				<rule avp="Framed-MTU" required="false" max="1"/>
				<rule avp="Framed-Protocol" required="false" max="1"/>
				<rule avp="Framed-Routing" required="false" max="1"/>
				<rule avp="Login-IP-Host" required="false"/>
				<rule avp="Login-IPv6-Host" required="false"/>
				<rule avp="Login-LAT-Group" required="false" max="1"/>
				<rule avp="Login-LAT-Node" required="false" max="1"/>
				<rule avp="Login-LAT-Port" required="false" max="1"/>
				<rule avp="Login-LAT-Service" required="false" max="1"/>
				<rule avp="Login-Service" required="false" max="1"/>
				<rule avp="Login-TCP-Port" required="false" max="1"/>
				<rule avp="NAS-Filter-Rule" required="false"/>
				<rule avp="QoS-Filter-Rule" required="false"/>
				<rule avp="Tunneling" required="false"/>
				<rule avp="Redirect-Host" required="false"/>
				<rule avp="Redirect-Host-Usage" required="false" max="1"/>
				<rule avp="Redirect-Max-Cache-Time" required="false" max="1"/>
				<rule avp="Proxy-Info" required="false"/>
			</answer>
		</command>

		<command code="258" short="RA" name="Re-Auth">
			<request>
				<!-- http://tools.ietf.org/html/rfc7155#section-3.3 -->
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Destination-Realm" required="true" max="1"/>
				<rule avp="Destination-Host" required="true" max="1"/>
				<rule avp="Auth-Application-Id" required="true" max="1"/>
				<rule avp="Re-Auth-Request-Type" required="true" max="1"/>
				<rule avp="User-Name" required="false" max="1"/>
				<rule avp="Origin-AAA-Protocol" required="false" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="NAS-Identifier" required="false" max="1"/>
				<rule avp="NAS-IP-Address" required="true" max="1"/>
				<rule avp="NAS-IPv6-Address" required="false" max="1"/>
				<rule avp="NAS-Port" required="false" max="1"/>
				<rule avp="NAS-Port-Id" required="false" max="1"/>
				<rule avp="NAS-Port-Type" required="false" max="1"/>
				<rule avp="Service-Type" required="false" max="1"/>
				<rule avp="Framed-IP-Address" required="false" max="1"/>
				<rule avp="Framed-IPv6-Prefix" required="false" max="1"/>
				<rule avp="Framed-Interface-Id" required="false" max="1"/>
				<rule avp="Called-Station-Id" required="false" max="1"/>
				<rule avp="Calling-Station-Id" required="false" max="1"/>
				<rule avp="Originating-Line-Info" required="false" max="1"/>
				<rule avp="Acct-Session-Id" required="false" max="1"/>
				<rule avp="Acct-Multi-Session-Id" required="false" max="1"/>
				<rule avp="State" required="false" max="1"/>
				<rule avp="Class" required="false"/>
				<rule avp="Reply-Message" required="false" max="1"/>
				<rule avp="Proxy-Info" required="false"/>
				<rule avp="Route-Record" required="false"/>
			</request>
			<answer>
				<!-- http://tools.ietf.org/html/rfc7155#section-3.4 -->
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Result-Code" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="User-Name" required="false" max="1"/>
				<rule avp="Origin-AAA-Protocol" required="false" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Error-Message" required="false" max="1"/>
				<rule avp="Error-Reporting-Host" required="false" max="1"/>
				<rule avp="Failed-AVP" required="false"/>
				<rule avp="Redirect-Host" required="false"/>
				<rule avp="Redirect-Host-Usage" required="false" max="1"/>
				<rule avp="Redirect-Max-Cache-Time" required="false" max="1"/>
				<rule avp="Service-Type" required="false" max="1"/>
				<rule avp="Configuration-Token" required="false"/>
				<rule avp="Idle-Timeout" required="false" max="1"/>
				<rule avp="Authorization-Lifetime" required="false" max="1"/>
				<rule avp="Auth-Grace-Period" required="false" max="1"/>
				<rule avp="Re-Auth-Request-Type" required="false" max="1"/>
				<rule avp="State" required="false" max="1"/>
				<rule avp="Class" required="false"/>
				<rule avp="Reply-Message" required="false"/>
				<rule avp="Prompt" required="false" max="1"/>
				<rule avp="Proxy-Info" required="false"/>
			</answer>
		</command>

		<command code="275" short="ST" name="Session-Termination">
			<request>
				<!-- http://tools.ietf.org/html/rfc7155#section-3.5 -->
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Destination-Realm" required="true" max="1"/>
				<rule avp="Auth-Application-Id" required="true" max="1"/>
				<rule avp="Termination-Cause" required="true" max="1"/>
				<rule avp="User-Name" required="false" max="1"/>
				<rule avp="Destination-Host" required="false" max="1"/>
				<rule avp="Class" required="false"/>
				<rule avp="Origin-AAA-Protocol" required="false" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Proxy-Info" required="false"/>
				<rule avp="Route-Record" required="false"/>
			</request>
			<answer>
				<!-- http://tools.ietf.org/html/rfc7155#section-3.6 -->
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Result-Code" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="User-Name" required="false" max="1"/>
				<rule avp="Class" required="false"/>
				<rule avp="Error-Message" required="false" max="1"/>
				<rule avp="Error-Reporting-Host" required="false" max="1"/>
				<rule avp="Failed-AVP" required="false"/>
				<rule avp="Origin-AAA-Protocol" required="false" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Redirect-Host" required="false"/>
				<rule avp="Redirect-Host-Usage" required="false" max="1"/>
				<rule avp="Redirect-Max-Cache-Time" required="false" max="1"/>
				<rule avp="Proxy-Info" required="false"/>
			</answer>
		</command>

		<command code="274" short="AS" name="Abort-Session">
			<request>
				<!-- http://tools.ietf.org/html/rfc7155#section-3.7 -->
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Destination-Realm" required="true" max="1"/>
				<rule avp="Destination-Host" required="true" max="1"/>
				<rule avp="Auth-Application-Id" required="true" max="1"/>
				<rule avp="User-Name" required="false" max="1"/>
				<rule avp="Origin-AAA-Protocol" required="false" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="NAS-Identifier" required="false" max="1"/>
				<rule avp="NAS-IP-Address" required="true" max="1"/>
				<rule avp="NAS-IPv6-Address" required="false" max="1"/>
				<rule avp="NAS-Port" required="false" max="1"/>
				<rule avp="NAS-Port-Id" required="false" max="1"/>
				<rule avp="NAS-Port-Type" required="false" max="1"/>
				<rule avp="Service-Type" required="false" max="1"/>
				<rule avp="Framed-IP-Address" required="false" max="1"/>
				<rule avp="Framed-IPv6-Prefix" required="false" max="1"/>
				<rule avp="Framed-Interface-Id" required="false" max="1"/>
				<rule avp="Called-Station-Id" required="false" max="1"/>
				<rule avp="Calling-Station-Id" required="false" max="1"/>
				<rule avp="Originating-Line-Info" required="false" max="1"/>
				<rule avp="Acct-Session-Id" required="false" max="1"/>
				<rule avp="Acct-Multi-Session-Id" required="false" max="1"/>
				<rule avp="State" required="false" max="1"/>
				<rule avp="Class" required="false"/>
				<rule avp="Reply-Message" required="false"/>
				<rule avp="Proxy-Info" required="false"/>
				<rule avp="Route-Record" required="false"/>
			</request>
			<answer>
				<!-- http://tools.ietf.org/html/rfc7155#section-3.8 -->
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Result-Code" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="User-Name" required="false" max="1"/>
				<rule avp="Origin-AAA-Protocol" required="false" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="State" required="false" max="1"/>
				<rule avp="Error-Message" required="false" max="1"/>
				<rule avp="Error-Reporting-Host" required="false" max="1"/>
				<rule avp="Failed-AVP" required="false"/>
				<rule avp="Redirect-Host" required="false"/>
				<rule avp="Redirect-Host-Usage" required="false" max="1"/>
				<rule avp="Redirect-Max-Cache-Time" required="false" max="1"/>
				<rule avp="Proxy-Info" required="false"/>
			</answer>
		</command>

		<command code="271" short="AC" name="Accounting">
			<request>
				<!-- http://tools.ietf.org/html/rfc7155#section-3.9 -->
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Destination-Realm" required="true" max="1"/>
				<rule avp="Accounting-Record-Type" required="true" max="1"/>
				<rule avp="Accounting-Record-Number" required="true" max="1"/>
				<rule avp="Acct-Application-Id" required="true" max="1"/>
				<rule avp="User-Name" required="false" max="1"/>
				<rule avp="Accounting-Sub-Session-Id" required="false" max="1"/>
				<rule avp="Acct-Session-Id" required="false" max="1"/>
				<rule avp="Acct-Multi-Session-Id" required="false" max="1"/>
				<rule avp="Origin-AAA-Protocol" required="false" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Destination-Host" required="false" max="1"/>
				<rule avp="Event-Timestamp" required="false" max="1"/>
				<rule avp="Acct-Delay-Time" required="false" max="1"/>
				<rule avp="NAS-Identifier" required="false" max="1"/>
				<rule avp="NAS-IP-Address" required="true" max="1"/>
				<rule avp="NAS-IPv6-Address" required="false" max="1"/>
				<rule avp="NAS-Port" required="false" max="1"/>
				<rule avp="NAS-Port-Id" required="false" max="1"/>
				<rule avp="NAS-Port-Type" required="false" max="1"/>
				<rule avp="Class" required="false"/>
				<rule avp="Service-Type" required="false" max="1"/>
				<rule avp="Termination-Cause" required="false" max="1"/>
				<rule avp="Accounting-Input-Octets" required="false" max="1"/>
				<rule avp="Accounting-Input-Packets" required="false" max="1"/>
				<rule avp="Accounting-Output-Octets" required="false" max="1"/>
				<rule avp="Accounting-Output-Packets" required="false" max="1"/>
				<rule avp="Acct-Authentic" required="false" max="1"/>
				<rule avp="Accounting-Auth-Method" required="false" max="1"/>
				<rule avp="Acct-Link-Count" required="false" max="1"/>
				<rule avp="Acct-Session-Time" required="false" max="1"/>
				<rule avp="Acct-Tunnel-Connection" required="false" max="1"/>
				<rule avp="Acct-Tunnel-Packets-Lost" required="false" max="1"/>
				<rule avp="Callback-Id" required="false" max="1"/>
				<rule avp="Callback-Number" required="false" max="1"/>
				<rule avp="Called-Station-Id" required="false" max="1"/>
				<rule avp="Calling-Station-Id" required="false" max="1"/>
				<rule avp="Connection-Info" required="false"/>
				<rule avp="Originating-Line-Info" required="false" max="1"/>
				<rule avp="Authorization-Lifetime" required="false" max="1"/>
				<rule avp="Session-Timeout" required="false" max="1"/>
				<rule avp="Idle-Timeout" required="false" max="1"/>
				<rule avp="Port-Limit" required="false" max="1"/>
				<rule avp="Accounting-Realtime-Required" required="false" max="1"/>
				<rule avp="Acct-Interim-Interval" required="false" max="1"/>
				<rule avp="Filter-Id" required="false"/>
				<rule avp="NAS-Filter-Rule" required="false"/>
				<rule avp="QoS-Filter-Rule" required="false"/>
				<rule avp="Framed-Appletalk-Link" required="false" max="1"/>
				<rule avp="Framed-Appletalk-Network" required="false" max="1"/>
				<rule avp="Framed-Appletalk-Zone" required="false" max="1"/>
				<rule avp="Framed-Compression" required="false" max="1"/>
				<rule avp="Framed-Interface-Id" required="false" max="1"/>
				<rule avp="Framed-IP-Address" required="false" max="1"/>
				<rule avp="Framed-IP-Netmask" required="false" max="1"/>
				<rule avp="Framed-IPv6-Prefix" required="false"/>
				<rule avp="Framed-IPv6-Pool" required="false" max="1"/>
				<rule avp="Framed-IPv6-Route" required="false"/>
				<rule avp="Framed-IPX-Network" required="false" max="1"/>
				<rule avp="Framed-MTU" required="false" max="1"/>
				<rule avp="Framed-Pool" required="false" max="1"/>
				<rule avp="Framed-Protocol" required="false" max="1"/>
				<rule avp="Framed-Route" required="false"/>
				<rule avp="Framed-Routing" required="false" max="1"/>
				<rule avp="Login-IP-Host" required="false"/>
				<rule avp="Login-IPv6-Host" required="false"/>
				<rule avp="Login-LAT-Group" required="false" max="1"/>
				<rule avp="Login-LAT-Node" required="false" max="1"/>
				<rule avp="Login-LAT-Port" required="false" max="1"/>
				<rule avp="Login-LAT-Service" required="false" max="1"/>
				<rule avp="Login-Service" required="false" max="1"/>
				<rule avp="Login-TCP-Port" required="false" max="1"/>
				<rule avp="Tunneling" required="false"/>
				<rule avp="Proxy-Info" required="false"/>
				<rule avp="Route-Record" required="false"/>
			</request>
			<answer>
				<!-- http://tools.ietf.org/html/rfc7155#section-3.10 -->
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Result-Code" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Accounting-Record-Type" required="true" max="1"/>
				<rule avp="Accounting-Record-Number" required="true" max="1"/>
				<rule avp="Acct-Application-Id" required="true" max="1"/>
				<rule avp="User-Name" required="false" max="1"/>
				<rule avp="Accounting-Sub-Session-Id" required="false" max="1"/>
				<rule avp="Acct-Session-Id" required="false" max="1"/>
				<rule avp="Acct-Multi-Session-Id" required="false" max="1"/>
				<rule avp="Event-Timestamp" required="false" max="1"/>
				<rule avp="Error-Message" required="false" max="1"/>
				<rule avp="Error-Reporting-Host" required="false" max="1"/>
				<rule avp="Failed-AVP" required="false"/>
				<rule avp="Origin-AAA-Protocol" required="false" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="NAS-Identifier" required="false" max="1"/>
				<rule avp="NAS-IP-Address" required="true" max="1"/>
				<rule avp="NAS-IPv6-Address" required="false" max="1"/>
				<rule avp="NAS-Port" required="false" max="1"/>
				<rule avp="NAS-Port-Id" required="false" max="1"/>
				<rule avp="NAS-Port-Type" required="false" max="1"/>
				<rule avp="Service-Type" required="false" max="1"/>
				<rule avp="Termination-Cause" required="false" max="1"/>
				<rule avp="Accounting-Realtime-Required" required="false" max="1"/>
				<rule avp="Acct-Interim-Interval" required="false" max="1"/>
				<rule avp="Class" required="false"/>
				<rule avp="Proxy-Info" required="false"/>
			</answer>
		</command>



		<avp name="NAS-Port" code="5" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.2.2 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="NAS-Port-Id" code="87" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.2.3 -->
			<data type="UTF8String"/>
		</avp>

		<avp name="NAS-Port-Type" code="61" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.2.4 -->
			<data type="Enumerated">
				<!-- http://www.iana.org/assignments/radius-types/radius-types.xhtml#radius-types-13 -->
				<item code="0" name="Async"/>
				<item code="1" name="Sync"/>
				<item code="2" name="ISDN Sync"/>
				<item code="3" name="ISDN Async V.120"/>
				<item code="4" name="ISDN Async V.110"/>
				<item code="5" name="Virtual"/>
				<item code="6" name="PIAFS"/>
				<item code="7" name="HDLC Clear Channel"/>
				<item code="8" name="X.25"/>
				<item code="9" name="X.75"/>
				<item code="10" name="G.3 Fax"/>
				<item code="11" name="SDSL - Symmetric DSL"/>
				<item code="12" name="ADSL-CAP - Asymmetric DSL, Carrierless Amplitude Phase Modulation"/>
				<item code="13" name="ADSL-DMT - Asymmetric DSL, Discrete Multi-Tone"/>
				<item code="14" name="IDSL - ISDN Digital Subscriber Line"/>
				<item code="15" name="Ethernet"/>
				<item code="16" name="xDSL - Digital Subscriber Line of unknown type"/>
				<item code="17" name="Cable"/>
				<item code="18" name="Wireless - Other"/>
				<item code="19" name="Wireless - IEEE 802.11"/>
				<item code="20" name="Token-Ring"/>
				<item code="21" name="FDDI"/>
				<item code="22" name="Wireless - CDMA2000"/>
				<item code="23" name="Wireless - UMTS"/>
				<item code="24" name="Wireless - 1X-EV"/>
				<item code="25" name="IAPP"/>
				<item code="26" name="FTTP - Fiber to the Premises"/>
				<item code="27" name="Wireless - IEEE 802.16"/>
				<item code="28" name="Wireless - IEEE 802.20"/>
				<item code="29" name="Wireless - IEEE 802.22"/>
				<item code="30" name="PPPoA - PPP over ATM"/>
				<item code="31" name="PPPoEoA - PPP over Ethernet over ATM"/>
				<item code="32" name="PPPoEoE - PPP over Ethernet over Ethernet"/>
				<item code="33" name="PPPoEoVLAN - PPP over Ethernet over VLAN"/>
				<item code="34" name="PPPoEoQinQ - PPP over Ethernet over IEEE 802.1QinQ"/>
				<item code="35" name="xPON - Passive Optical Network"/>
				<item code="36" name="Wireless - XGP"/>
				<item code="37" name="WiMAX Pre-Release 8 IWK Function"/>
				<item code="38" name="WIMAX-WIFI-IWK: WiMAX WIFI Interworking"/>
				<item code="39" name="WIMAX-SFF: Signaling Forwarding Function for LTE/3GPP2"/>
				<item code="40" name="WIMAX-HA-LMA: WiMAX HA and or LMA function"/>
				<item code="41" name="WIMAX-DHCP: WIMAX DCHP service"/>
				<item code="42" name="WIMAX-LBS: WiMAX location based service"/>
				<item code="43" name="WIMAX-WVS: WiMAX voice service"/>
			</data>
		</avp>

		<avp name="Called-Station-Id" code="30" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.2.5 -->
			<data type="UTF8String"/>
		</avp>

		<avp name="Calling-Station-Id" code="31" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.2.6 -->
			<data type="UTF8String"/>
		</avp>

		<avp name="Connect-Info" code="77" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.2.7 -->
			<data type="UTF8String"/>
		</avp>

		<avp name="Originating-Line-Info" code="94" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.2.8 -->
			<data type="OctetString"/>
		</avp>

		<avp name="Reply-Message" code="18" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.2.9 -->
			<data type="UTF8String"/>
		</avp>

		<avp name="User-Password" code="2" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.3.1 -->
			<data type="OctetString"/>
		</avp>

		<avp name="Password-Retry" code="75" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.3.2 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="Prompt" code="76" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.3.3 -->
			<data type="Enumerated">
				<!-- http://www.iana.org/assignments/radius-types/radius-types.xhtml#radius-types-17 -->
				<item code="0" name="No Echo"/>
				<item code="1" name="Echo"/>
			</data>
		</avp>

		<avp name="CHAP-Auth" code="402" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.3.4 -->
			<data type="Grouped">
				<rule avp="CHAP-Algorithm" required="true" max="1"/>
				<rule avp="CHAP-Ident" required="true" max="1"/>
				<rule avp="CHAP-Response" required="true" max="1"/>
			</data>
		</avp>


		<avp name="CHAP-Algorithm" code="403" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.3.5 -->
			<data type="Enumerated">
				<item code="5" name="CHAP with MD5"/>
			</data>
		</avp>

		<avp name="CHAP-Ident" code="404" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.3.6 -->
			<data type="OctetString"/>
		</avp>

		<avp name="CHAP-Response" code="405" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.3.7 -->
			<data type="OctetString"/>
		</avp>

		<avp name="CHAP-Challenge" code="60" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.3.8 -->
			<data type="OctetString"/>
		</avp>

		<avp name="ARAP-Password" code="70" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.3.9 -->
			<data type="OctetString"/>
		</avp>

		<avp name="ARAP-Challenge-Response" code="84" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.3.10 -->
			<data type="OctetString"/>
		</avp>

		<avp name="ARAP-Security" code="73" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.3.11 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="ARAP-Security-Data" code="74" must="M" may="" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.3.12 -->
			<data type="OctetString"/>
		</avp>

		<avp name="Service-Type" code="6" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.4.1 -->
			<data type="Enumerated">
				<!-- http://www.iana.org/assignments/radius-types/radius-types.xhtml#radius-types-4 -->
				<item code="1" name="Login"/>
				<item code="2" name="Framed"/>
				<item code="3" name="Callback Login"/>
				<item code="4" name="Callback Framed"/>
				<item code="5" name="Outbound"/>
				<item code="6" name="Administrative"/>
				<item code="7" name="NAS Prompt"/>
				<item code="8" name="Authenticate Only"/>
				<item code="9" name="Callback NAS Prompt"/>
				<item code="10" name="Call Check"/>
				<item code="11" name="Callback Administrative"/>
				<item code="12" name="Voice"/>
				<item code="13" name="Fax"/>
				<item code="14" name="Modem Relay"/>
				<item code="15" name="IAPP-Register"/>
				<item code="16" name="IAPP-AP-Check"/>
				<item code="17" name="Authorize Only"/>
				<item code="18" name="Framed-Management"/>
				<item code="19" name="Additional-Authorization"/>
			</data>
		</avp>

		<avp name="Callback-Number" code="19" must="M" may="" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.4.2 -->
			<data type="UTF8String"/>
		</avp>

		<avp name="Callback-Id" code="20" must="M" may="" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.4.3 -->
			<data type="UTF8String"/>
		</avp>

		<avp name="Idle-Timeout" code="28" must="M" may="" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.4.4 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="Port-Limit" code="62" must="M" may="" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.4.5 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="NAS-Filter-Rule" code="400" must="M" may="" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.4.6 -->
			<data type="IPFilterRule"/>
		</avp>

		<avp name="Filter-Id" code="11" must="M" may="" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.4.7 -->
			<data type="UTF8String"/>
		</avp>

		<avp name="Configuration-Token" code="78" must="M" may="" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.4.8 -->
			<data type="OctetString"/>
		</avp>

		<!--avp name="QoS-Filter-Rule" code="407" must="-" may="" must-not="-" may-encrypt="Y"-->
			<!-- http://tools.ietf.org/html/rfc7155#section-4.4.9 -->
			<!--data type="QoSFilterRule"/-->
		<!--/avp-->


		<avp name="Framed-Protocol" code="7" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.4.10.1 -->
			<data type="Enumerated">
				<!-- http://www.iana.org/assignments/radius-types/radius-types.xhtml#radius-types-5 -->
				<item code="1" name="PPP"/>
				<item code="2" name="SLIP"/>
				<item code="3" name="AppleTalk Remote Access Protocol (ARAP)"/>
				<item code="4" name="Gandalf proprietary SingleLink/MultiLink protocol	"/>
				<item code="5" name="Xylogics proprietary IPX/SLIP"/>
				<item code="6" name="X.75 Synchronous"/>
				<item code="7" name="GPRS PDP Context"/>
			</data>
		</avp>

		<avp name="Framed-Routing" code="10" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.4.10.2 -->
			<data type="Enumerated">
				<!-- http://www.iana.org/assignments/radius-types/radius-types.xhtml#radius-types-6 -->
				<item code="0" name="None"/>
				<item code="1" name="Send routing packets"/>
				<item code="2" name="Listen for routing packets"/>
				<item code="3" name="Send and Listen"/>
			</data>
		</avp>

		<avp name="Framed-MTU" code="12" must="M" may="" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.4.10.3 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="Framed-Compression" code="13" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.4.10.4 -->
			<data type="Enumerated">
				<!-- http://www.iana.org/assignments/radius-types/radius-types.xhtml#radius-types-7 -->
				<item code="0" name="None"/>
				<item code="1" name="VJ TCP/IP header compression	"/>
				<item code="2" name="IPX header compression"/>
				<item code="3" name="Stac-LZS compression"/>
			</data>
		</avp>

		<avp name="Framed-IP-Address" code="8" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.4.10.5.1 -->
			<data type="OctetString"/>
		</avp>

		<avp name="Framed-IP-Netmask" code="9" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.4.10.5.2 -->
			<data type="OctetString"/>
		</avp>

		<avp name="Framed-Route" code="22" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.4.10.5.3 -->
			<data type="UTF8String"/>
		</avp>

		<avp name="Framed-Pool" code="88" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.4.10.5.4 -->
			<data type="OctetString"/>
		</avp>

		<avp name="Framed-Interface-Id" code="96" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.4.10.5.5 -->
			<data type="Unsigned64"/>
		</avp>

		<avp name="Framed-IPv6-Prefix" code="97" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.4.10.5.6 -->
			<data type="OctetString"/>
		</avp>

		<avp name="Framed-IPv6-Route" code="99" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.4.10.5.7 -->
			<data type="UTF8String"/>
		</avp>

		<avp name="Framed-IPv6-Pool" code="100" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.4.10.5.8 -->
			<data type="OctetString"/>
		</avp>

		<avp name="Framed-IPX-Network" code="23" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.4.10.6.1-->
			<data type="Unsigned32"/>
		</avp>

		<avp name="Framed-Appletalk-Link" code="37" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.4.10.7.1-->
			<data type="Unsigned32"/>
		</avp>

		<avp name="Framed-Appletalk-Network" code="38" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.4.10.7.2-->
			<data type="Unsigned32"/>
		</avp>

		<avp name="Framed-Appletalk-Zone" code="39" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.4.10.7.3-->
			<data type="OctetString"/>
		</avp>

		<avp name="ARAP-Features" code="71" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.4.10.8.1-->
			<data type="OctetString"/>
		</avp>

		<avp name="ARAP-Zone-Access" code="72" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.4.10.8.2 -->
			<data type="Enumerated">
				<!-- http://www.iana.org/assignments/radius-types/radius-types.xhtml#radius-types-16 -->
				<item code="1" name="Only allow access to default zone"/>
				<item code="2" name="Use zone filter inclusively"/>
				<item code="3" name="Not used"/>
				<item code="4" name="Use zone filter exclusively"/>
			</data>
		</avp>

		<avp name="Login-IP-Host" code="14" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.4.11.1-->
			<data type="OctetString"/>
		</avp>

		<avp name="Login-IPv6-Host" code="98" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.4.11.2-->
			<data type="OctetString"/>
		</avp>

		<avp name="Login-Service" code="15" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.4.11.3 -->
			<data type="Enumerated">
				<!-- http://www.iana.org/assignments/radius-types/radius-types.xhtml#radius-types-8 -->
				<item code="0" name="Telnet"/>
				<item code="1" name="Rlogin"/>
				<item code="2" name="TCP Clear"/>
				<item code="3" name="PortMaster (proprietary)"/>
				<item code="4" name="LAT"/>
				<item code="5" name="X25-PAD"/>
				<item code="6" name="X25-T3POS"/>
				<item code="7" name="Unassigned"/>
				<item code="8" name="TCP Clear Quiet (suppresses any NAS-generated connect string)"/>
			</data>
		</avp>

		<avp name="Login-TCP-Port" code="16" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.4.11.4.1-->
			<data type="Unsigned32"/>
		</avp>

		<avp name="Login-LAT-Service" code="34" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.4.11.5.1-->
			<data type="OctetString"/>
		</avp>

		<avp name="Login-LAT-Node" code="35" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.4.11.5.2-->
			<data type="OctetString"/>
		</avp>

		<avp name="Login-LAT-Group" code="36" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.4.11.5.3-->
			<data type="OctetString"/>
		</avp>

		<avp name="Login-LAT-Port" code="63" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.4.11.5.4-->
			<data type="OctetString"/>
		</avp>

		<avp name="Tunneling" code="401" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.5.1-->
			<data type="Grouped">
				<rule avp="Tunnel-Type" required="true" max="1"/>
				<rule avp="Tunnel-Medium-Type" required="true" max="1"/>
				<rule avp="Tunnel-Client-Endpoint" required="true" max="1"/>
				<rule avp="Tunnel-Server-Endpoint" required="true" max="1"/>
				<rule avp="Tunnel-Preference" required="false" max="1"/>
				<rule avp="Tunnel-Client-Auth-Id" required="false" max="1"/>
				<rule avp="Tunnel-Server-Auth-Id" required="false" max="1"/>
				<rule avp="Tunnel-Assignment-Id" required="false" max="1"/>
				<rule avp="Tunnel-Password" required="false" max="1"/>
				<rule avp="Tunnel-Private-Group-Id" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Tunnel-Type" code="64" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.5.2 -->
			<data type="Enumerated">
				<!-- http://www.iana.org/assignments/radius-types/radius-types.xhtml#radius-types-14 -->
				<item code="1" name="Point-to-Point Tunneling Protocol (PPTP)"/>
				<item code="2" name="Layer Two Forwarding (L2F)"/>
				<item code="3" name="Layer Two Tunneling Protocol (L2TP)"/>
				<item code="4" name="Ascend Tunnel Management Protocol (ATMP)"/>
				<item code="5" name="Virtual Tunneling Protocol (VTP)"/>
				<item code="6" name="IP Authentication Header in the Tunnel-mode (AH)"/>
				<item code="7" name="IP-in-IP Encapsulation (IP-IP)"/>
				<item code="8" name="Minimal IP-in-IP Encapsulation (MIN-IP-IP)"/>
				<item code="9" name="IP Encapsulating Security Payload in the Tunnel-mode (ESP)"/>
				<item code="10" name="Generic Route Encapsulation (GRE)"/>
				<item code="11" name="Bay Dial Virtual Services (DVS)"/>
				<item code="12" name="IP-in-IP Tunneling"/>
				<item code="13" name="Virtual LANs (VLAN)"/>
			</data>
		</avp>

		<avp name="Tunnel-Medium-Type" code="65" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.5.3 -->
			<data type="Enumerated">
				<!-- http://www.iana.org/assignments/radius-types/radius-types.xhtml#radius-types-15 -->
				<item code="1" name="IPv4 (IP version 4)"/>
				<item code="2" name="IPv6 (IP version 6)"/>
				<item code="3" name="NSAP"/>
				<item code="4" name="HDLC (8-bit multidrop)"/>
				<item code="5" name="BBN 1822"/>
				<item code="6" name="802 (includes all 802 media plus Ethernet 'canonical format')"/>
				<item code="7" name="E.163 (POTS)"/>
				<item code="8" name="E.164 (SMDS, Frame Relay, ATM)"/>
				<item code="9" name="F.69 (Telex)"/>
				<item code="10" name="X.121 (X.25, Frame Relay)"/>
				<item code="11" name="IPX"/>
				<item code="12" name="Appletalk"/>
				<item code="13" name="Decnet IV"/>
				<item code="14" name="Banyan Vines"/>
				<item code="15" name="E.164 with NSAP format subaddress"/>
			</data>
		</avp>

		<avp name="Tunnel-Client-Endpoint" code="66" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.5.4 -->
			<data type="UTF8String"/>
		</avp>

		<avp name="Tunnel-Server-Endpoint" code="67" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.5.5 -->
			<data type="UTF8String"/>
		</avp>

		<avp name="Tunnel-Password" code="69" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.5.6 -->
			<data type="OctetString"/>
		</avp>

		<avp name="Tunnel-Private-Group-Id" code="81" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.5.7 -->
			<data type="OctetString"/>
		</avp>

		<avp name="Tunnel-Assignment-Id" code="82" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.5.8 -->
			<data type="OctetString"/>
		</avp>

		<avp name="Tunnel-Preference" code="83" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.5.9 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="Tunnel-Client-Auth-Id" code="90" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.5.10 -->
			<data type="UTF8String"/>
		</avp>

		<avp name="Tunnel-Server-Auth-Id" code="91" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.5.11 -->
			<data type="UTF8String"/>
		</avp>

		<avp name="Accounting-Input-Octets" code="363" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.6.1 -->
			<data type="Unsigned64"/>
		</avp>

		<avp name="Accounting-Output-Octets" code="364" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.6.2 -->
			<data type="Unsigned64"/>
		</avp>

		<avp name="Accounting-Input-Packets" code="365" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.6.3 -->
			<data type="Unsigned64"/>
		</avp>

		<avp name="Accounting-Output-Packets" code="366" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.6.4 -->
			<data type="Unsigned64"/>
		</avp>

		<avp name="Acct-Session-Time" code="46" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.6.5 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="Acct-Authentic" code="45" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.6.6 -->
			<data type="Enumerated">
				<!-- http://www.iana.org/assignments/radius-types/radius-types.xhtml#radius-types-11 -->
				<item code="1" name="RADIUS"/>
				<item code="2" name="Local"/>
				<item code="3" name="Remote"/>
				<item code="4" name="Diameter"/>
			</data>
		</avp>

		<avp name="Accounting-Auth-Method" code="406" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.6.7 -->
			<data type="Enumerated">
				<!-- http://www.iana.org/assignments/aaa-parameters/aaa-parameters.xhtml#aaa-parameters-26 -->
				<item code="1" name="PAP"/>
				<item code="2" name="CHAP"/>
				<item code="3" name="MS-CHAP-1"/>
				<item code="4" name="MS-CHAP-2"/>
				<item code="5" name="EAP"/>
				<item code="7" name="None"/>
			</data>
		</avp>

		<avp name="Acct-Delay-Time" code="41" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.6.8 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="Acct-Link-Count" code="51" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.6.9 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="Acct-Tunnel-Connection" code="68" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.6.10 -->
			<data type="OctetString"/>
		</avp>

		<avp name="Acct-Tunnel-Packets-Lost" code="86" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc7155#section-4.6.11 -->
			<data type="Unsigned32"/>
		</avp>
		
	</application>
</diameter>`

var tgppe2XML = `<?xml version="1.0" encoding="UTF-8"?>
<diameter>
        <!--
        3GPP TS 283 035
        See: https://www.etsi.org/deliver/etsi_es/283000_283099/283035/03.02.01_60/es_283035v030201p.pdf
    -->
    <application id="16777231" type="auth" name="TGPP e2">
         <vendor id="10415" name="TGPP"/>
         <command code="306" short="UD" name="User-Data">
            <!--
                < User-Data -Request > ::= < Diameter Header: 306, REQ, PXY, 16777231 >
                < Session-Id >
                { Vendor-Specific-Application-Id }
                { Auth-Session-State }
                { Origin-Host }
                { Origin-Realm }
                [ Destination-Host ]
                { Destination-Realm }
                [Globally-Unique-Address]
                [User-Name]
                [Global-Access-Id]
                [AF-Application-Identifier]
                [Requested-Information]
                *[ AVP ]
                *[ Proxy-Info ]
                *[ Route-Record ]
            -->
            <request>
                <rule avp="Session-Id" required="true" max="1" />
                <rule avp="Vendor-Specific-Application-Id" required="true" max="1" />
                <rule avp="Auth-Session-State" required="true" max="1" />
                <rule avp="Origin-Host" required="true" max="1" />
                <rule avp="Origin-Realm" required="true" max="1" />
                <rule avp="Destination-Host" required="false" max="1" />
                <rule avp="Destination-Realm" required="true" max="1" />
                <rule avp="Globally-Unique-Address" required="false" max="1" />
                <rule avp="User-Name" required="false" max="1" />
                <rule avp="Global-Access-Id" required="false" max="1" />
                <rule avp="AF-Application-Identifier" required="false" max="1" />
                <rule avp="Requested-Information" required="false" max="1" />
                <rule avp="Proxy-Info" required="false" />
                <rule avp="Route-Record" required="false" />                
            </request>

            <!--
                < User-Data-Answer > ::= < Diameter Header: 306, PXY, 16777231 >
                < Session-Id >
                { Vendor-Specific-Application-Id }
                [ Result-Code ]
                [ Experimental-Result ]
                { Auth-Session-State }
                { Origin-Host }
                { Origin-Realm }
                [User-Name]
                [Logical-Access-Id]
                [Physical-Access-Id]
                [Access-Network-Type]
                [Location-Information]
                [Policy-Control-Contact-Point]
                [Terminal-Type]
                *[ AVP ]
                *[ Failed-AVP ]
                *[ Proxy-Info ]
                *[ Route-Record ]
            -->
            <answer>
                <rule avp="Session-Id" required="true" max="1" />
                <rule avp="Vendor-Specific-Application-Id" required="true" max="1" />
                <rule avp="Result-Code" required="false" max="1" />
                <rule avp="Experimental-Result" required="false" max="1" />
                <rule avp="Auth-Session-State" required="true" max="1" />
                <rule avp="Origin-Host" required="true" max="1" />
                <rule avp="Origin-Realm" required="true" max="1" />
                <rule avp="User-Name" required="false" max="1" />
                <rule avp="Logical-Access-Id" required="false" max="1" />
                <rule avp="Physical-Access-Id" required="false" max="1" />
                <rule avp="AF-Application-Identifier" required="false" max="1" />
                <rule avp="Location-Information" required="false" max="1" />
                <rule avp="Policy-Control-Contact-Point" required="false" max="1" />
                <rule avp="Terminal-Type" required="false" max="1" />
                <rule avp="Failed-AVP" required="false" max="1" />
                <rule avp="Proxy-Info" required="false" />
                <rule avp="Route-Record" required="false" />                                
            </answer>
         </command>

         <command code="308" short="SN" name="Subscribe-Notifications">

            <!--
                < Subscribe-Notifications-Request > ::= < Diameter Header: 308, REQ, PXY, 16777231 >
                < Session-Id >
                { Vendor-Specific-Application-Id }
                { Auth-Session-State }
                { Origin-Host }
                { Origin-Realm }
                [ Destination-Host ]
                { Destination-Realm }
                { Subs-Req-Type }
                [ Expiry-Time ]
                [Globally-Unique-Address]
                [User-Name]
                [AF-Application-Identifier]
                *[Event-Type]
                *[ AVP ]
                *[ Proxy-Info ]
                *[ Route-Record ]           
            -->
            <request>
            <!-- TBD -->
            </request>

            <!--
                < Subscribe-Notifications-Answer > ::= < Diameter Header: 308, PXY, 16777231 >
                < Session-Id >
                { Vendor-Specific-Application-Id }
                { Auth-Session-State }
                [ Result-Code ]
                [ Experimental-Result ]
                { Origin-Host }
                { Origin-Realm }
                [ Expiry-Time ]
                *[ AVP ]
                *[ Failed-AVP ]
                *[ Proxy-Info ]
                *[ Route-Record ]           
            -->
            <answer>
            <!-- TBD -->
            </answer>
         </command>

         <command code="309" short="PN" name="Push-Notification">

            <!--
                < Push-Notification-Request > ::= < Diameter Header: 309, REQ, PXY, 16777231 >
                < Session-Id >
                { Vendor-Specific-Application-Id }
                { Auth-Session-State }
                { Origin-Host }
                { Origin-Realm }
                { Destination-Host }
                { Destination-Realm }
                *[Event-Type]
                [Globally-Unique-Address]
                [User-Name]
                [Access-Network-Type]
                [Location-Information]
                [Policy-Control-Contact-Point]
                [Terminal-Type]
                [Logical-Access-Id]
                [Physical-Access-Id]
                [Access-Network-Type]
                [Initial-Gate-Setting]
                *[QoS-Profile]
                [IP-Connectivity-Status]
                *[ AVP ]
                *[ Proxy-Info ]
                *[ Route-Record ]           
            -->
            <request>
            <!-- TBD -->
            </request>

            <!--
                < Push-Notification-Answer > ::= < Diameter Header: 309, PXY, 16777231 >
                < Session-Id >
                { Vendor-Specific-Application-Id }
                [ Result-Code ]
                [ Experimental-Result ]
                { Auth-Session-State }
                { Origin-Host }
                { Origin-Realm }
                *[ AVP ]
                *[ Failed-AVP ]
                *[ Proxy-Info ]
                *[ Route-Record ]           
            -->
            <answer>
            <!-- TBD -->
            </answer>
         </command>

        <avp name="Location-Information" code="350" vendor-id="10415" must="V" may="M" may-encrypt="Y">
            <data type="Grouped">
                <rule avp="Line-Identifier" required="false" max="1" />
                <rule avp="Civic-Location" required="false" max="1" />
                <rule avp="Geospatial-Location" required="false" max="1" />
                <rule avp="Operator-Specific-GI" required="false" max="1" />
                <rule avp="PIDF-Location-Object" required="false" max="1" />
                <rule avp="AVP" required="false" />
            </data>
        </avp>

        <avp name="Policy-Control-Contact-Point" code="351" vendor-id="10415" must="V" may="M" may-encrypt="Y">
            <data type="DiameterIdentity"/>
        </avp>

        <avp name="Terminal-Type" code="352" vendor-id="10415" must="V" may="M" may-encrypt="Y">
            <data type="OctetString"/>
        </avp>
        
        <avp name="Requested-Information" code="353" vendor-id="10415" must="V" must-not="M" may-encrypt="Y">
            <data type="Enumerated">
                <item code="0" name="NASS-USER-ID"/>
                <item code="1" name="LOCATION-INFORMATION"/>
                <item code="2" name="POLICY-CONTROL-CONTACT-POINT"/>
                <item code="3" name="ACCESS-NETWORK-TYPE"/>
                <item code="4" name="TERMINAL-TYPE"/>
                <item code="5" name="LOGICAL-ACCESS-ID"/>
                <item code="6" name="PHYSICAL-ACCESS-ID"/>
                <item code="7" name="IP-ADDRESS"/>
                <item code="8" name="INITIAL-GATE-SETTING"/>
                <item code="9" name="QOS-PROFILE"/>
                <item code="10" name="IP-CONNECTIVITY-STATUS"/>
                <item code="11" name="EMERGENCY-CALL-ROUTING-INFO"/>
            </data>
        </avp>

        <avp name="Event-Type" code="354" vendor-id="10415" must="V" may="M" may-encrypt="Y">
            <data type="Enumerated">
                <item code="0" name="USER-LOGON"/>
                <item code="1" name="LOCATION-INFORMATION-CHANGED"/>
                <item code="2" name="POLICY-CONTROL-CONTACT-POINT-CHANGED"/>
                <item code="3" name="ACCESS-NETWORK-TYPE-CHANGED"/>
                <item code="4" name="TERMINAL-TYPE-CHANGED"/>
                <item code="5" name="LOGICAL-ACCESS-ID-CHANGED"/>
                <item code="6" name="PHYSICAL-ACCESS-ID-CHANGED"/>
                <item code="7" name="IP-ADDRESS-CHANGED"/>
                <item code="8" name="INITIAL-GATE-SETTING-CHANGED"/>
                <item code="9" name="QOS-PROFILE-CHANGED"/>
                <item code="10" name="USER-LOGOFF"/>
            </data>
        </avp>

        <avp name="Line-Identifier" code="500" vendor-id="10415" must="V" must-not="M" may-encrypt="Y">
            <data type="OctetString"/>
        </avp>

        <avp name="Civic-Location" code="355" vendor-id="10415" must="V" must-not="M" may-encrypt="Y">
            <data type="OctetString"/>
        </avp>

        <avp name="Geospatial-Location" code="356" vendor-id="10415" must="V" must-not="M" may-encrypt="Y">
            <data type="OctetString"/>
        </avp>

        <avp name="Global-Access-Id" code="357" vendor-id="10415" must="V" must-not="M" may-encrypt="Y">
            <data type="Grouped">
                <rule avp="Fixed-Access-ID" required="false" max="1" />
                <rule avp="3GPP-User-Location-Info" required="false" max="1" />
                <rule avp="AVP" required="false" />
            </data>
        </avp>

        <avp name="Fixed-Access-Id" code="358" vendor-id="10415" must="V" must-not="M" may-encrypt="Y">
            <data type="Grouped">
                <rule avp="Relay-Agent" required="true" max="1" />
                <rule avp="Logical-Access-ID" required="false" max="1" />
                <rule avp="Physical-Access-ID" required="false" max="1" />
                <rule avp="AVP" required="false" />
            </data>
        </avp>

        <avp name="Relay-Agent" code="359" vendor-id="10415" must="V" may="M" may-encrypt="Y">
            <data type="OctetString"/>
        </avp>

        <avp name="Operator-Specific-GI" code="360" vendor-id="10415" must="V" must-not="M" may-encrypt="Y">
            <data type="OctetString"/>
        </avp>

        <avp name="Emergency-Call-Routing-Info" code="361" vendor-id="10415" must="V" may="M" may-encrypt="Y">
            <data type="UTF8String"/>
        </avp>

        <avp name="Port-Number" code="362" vendor-id="10415" must="V" may="M" may-encrypt="Y">
            <data type="Unsigned32"/>
        </avp>

        <avp name="PIDF-Location-Object" code="363" vendor-id="10415" must="V" must-not="M" may-encrypt="Y">
            <data type="UTF8String"/>
        </avp>

        <avp name="AF-Application-Identifier" code="504" vendor-id="10415" must="M,V" may-encrypt="Y">
            <data type="OctetString"/>
        </avp>

        <avp name="Globally-Unique-Address" code="300" vendor-id="10415" must="M,V" may-encrypt="Y">
            <data type="Grouped">
                <rule avp="Frame-IP-Address" required="false" max="1" />
                <rule avp="Frame-IPv6-Prefix" required="false" max="1" />
                <rule avp="Address-Realm" required="false" max="1" />
            </data>
        </avp>

        <avp name="Logical-Access-Id" code="302" vendor-id="10415" must="V" may="M" may-encrypt="Y">
            <data type="OctetString"/>
        </avp>

        <avp name="Access-Network-Type" code="306" vendor-id="10415" must="V" may="M" may-encrypt="Y">
            <data type="Grouped">
                <rule avp="NAS-Port-Type" required="true" max="1" />
                <rule avp="Aggregation-Network-Type" required="false" max="1" />
            </data>
        </avp>

        <avp name="Initial-Gate-Setting" code="303" vendor-id="10415" must="V" may="M" may-encrypt="Y">
            <data type="Grouped">
                <rule avp="NAS-Filter-Rule" required="false" max="1" />
                <rule avp="Max-Requested-Bandwidth-UL" required="false" max="1" />
                <rule avp="Max-Requested-Bandwidth-DL" required="false" max="1" />
            </data>
        </avp>

        <avp name="QoS-Profile" code="304" vendor-id="10415" must="V" may="M" may-encrypt="Y">
            <data type="Grouped">
                <rule avp="Application-Class-ID" required="false" max="1" />
                <rule avp="Media-Type" required="false" max="1" />
                <rule avp="Reservation-Priority" required="false" max="1" />
                <rule avp="Max-Requested-Bandwidth-UL" required="false" max="1" />
                <rule avp="Max-Requested-Bandwidth-DL" required="false" max="1" />
                <rule avp="Transport-Class" required="false" max="1" />
            </data>
        </avp>

        <avp name="IP-Connectivity-Status" code="305" vendor-id="10415" must="V" may="M" may-encrypt="Y">
            <data type="Enumerated">
                <item code="0" name="IP-CONNECTIVITY-ON"/>
                <item code="1" name="IP-CONNECTIVITY-LOST"/>
            </data>
        </avp>

        <avp name="Aggregation-Network-Type" code="307" vendor-id="10415" must="V" may="M" may-encrypt="Y">
            <data type="Enumerated">
                <item code="0" name="UNKNOWN"/>
                <item code="1" name="ATM"/>
                <item code="2" name="ETHERNET"/>
            </data>
        </avp>

        <avp name="Expiry-Time" code="709" vendor-id="10415" must="V" must-not="M" may-encrypt="Y">
            <data type="Time">
            </data>
        </avp>

        <avp name="Subs-Req-Type" code="705" vendor-id="10415" must="M,V" may-encrypt="Y">
            <data type="Enumerated">
            <!-- TBD -->
            </data>
        </avp>
        
        <avp name="TGPP-User-Location-Info" code="22" vendor-id="10415" must="V" must-not="M" may-encrypt="Y">
            <data type="OctetString"/>
        </avp>

        <avp name="Frame-IP-Address" code="8" vendor-id="10415" may="M" must-not="V" may-encrypt="Y">
            <data type="OctetString"/>
        </avp>

        <avp name="Frame-IPv6-Prefix" code="97" vendor-id="10415" may="M" must-not="V" may-encrypt="Y">
            <data type="OctetString"/>
        </avp>
        
        <avp name="Max-Requested-Bandwidth-DL" code="515" must="V,M"    may="P" must-not="-" may-encrypt="Y" vendor-id="10415">
            <data type="Unsigned32"/>
        </avp>

        <avp name="Max-Requested-Bandwidth-UL" code="516" must="V,M"    may="P" must-not="-" may-encrypt="Y" vendor-id="10415">
            <data type="Unsigned32"/>
        </avp>


    </application>
</diameter>`

var tgpps6aXML = `<?xml version="1.0" encoding="UTF-8"?>
<diameter>
    <!--
        3GPP TS 29.272
        See: http://www.etsi.org/deliver/etsi_ts/129200_129299/129272/12.06.00_60/ts_129272v120600p.pdf
    -->
    <application id="16777251" type="auth" name="TGPP S6A">
        <vendor id="10415" name="TGPP"/>
        <command code="316" short="UL" name="Update-Location">
            <request>
                <rule avp="Session-Id" required="true" max="1"/>
                <rule avp="Vendor-Specific-Application-Id" required="false" max="1"/>
                <rule avp="Auth-Session-State" required="true" max="1"/>
                <rule avp="Origin-Host" required="true" max="1"/>
                <rule avp="Origin-Realm" required="true" max="1"/>
                <rule avp="Destination-Host" required="false" max="1"/>
                <rule avp="Destination-Realm" required="true" max="1"/>
                <rule avp="User-Name" required="true" max="1"/>
                <rule avp="Supported-Features" required="false"/>
                <rule avp="Terminal-Information" required="false" max="1"/>
                <rule avp="RAT-Type" required="true" max="1"/>
                <rule avp="ULR-Flags" required="true" max="1"/>
                <rule avp="UE-SRVCC-Capability" required="false" max="1"/>
                <rule avp="Visited-PLMN-Id" required="true" max="1"/>
                <rule avp="SGSN-Number" required="false" max="1"/>
                <rule avp="Homogeneous-Support-of-IMS-Voice-Over-PS-Sessions" required="false" max="1"/>
                <rule avp="GMLC-Address" required="false" max="1"/>
                <rule avp="Active-APN" required="false"/>
                <rule avp="AVP" required="false"/>
                <rule avp="Proxy-Info" required="false"/>
                <rule avp="Route-Record" required="false"/>
            </request>
            <answer>
                <rule avp="Session-Id" required="true" max="1"/>
                <rule avp="Vendor-Specific-Application-Id" required="false" max="1"/>
                <rule avp="Result-Code" required="false" max="1"/>
                <rule avp="Experimental-Result" required="false" max="1"/>
                <rule avp="Error-Diagnostic" required="false" max="1"/>
                <rule avp="Auth-Session-State" required="true" max="1"/>
                <rule avp="Origin-Host" required="true" max="1"/>
                <rule avp="Origin-Realm" required="true" max="1"/>
                <rule avp="Supported-Features" required="false"/>
                <rule avp="ULA-Flags" required="false" max="1"/>
                <rule avp="Subscription-Data" required="false" max="1"/>
                <rule avp="AVP" required="false"/>
                <rule avp="Failed-AVP" required="false"/>
                <rule avp="Proxy-Info" required="false"/>
                <rule avp="Route-Record" required="false"/>
            </answer>
        </command>

        <command code="317" short="CL" name="Cancel-Location">
            <request>
                <rule avp="Session-Id" required="true" max="1"/>
                <rule avp="Vendor-Specific-Application-Id" required="false" max="1"/>
                <rule avp="Auth-Session-State" required="true" max="1"/>
                <rule avp="Origin-Host" required="true" max="1"/>
                <rule avp="Origin-Realm" required="true" max="1"/>
                <rule avp="Destination-Host" required="true" max="1"/>
                <rule avp="Destination-Realm" required="true" max="1"/>
                <rule avp="User-Name" required="true" max="1"/>
                <rule avp="Supported-Features" required="false"/>
                <rule avp="Cancellation-Type" required="true" max="1"/>
                <rule avp="CLR-Flags" required="false" max="1"/>
                <rule avp="AVP" required="false"/>
                <rule avp="Proxy-Info" required="false"/>
                <rule avp="Route-Record" required="false"/>
            </request>
            <answer>
                <rule avp="Session-Id" required="true" max="1"/>
                <rule avp="Vendor-Specific-Application-Id" required="false" max="1"/>
                <rule avp="Supported-Features" required="false"/>
                <rule avp="Result-Code" required="false" max="1"/>
                <rule avp="Experimental-Result" required="false" max="1"/>
                <rule avp="Auth-Session-State" required="true" max="1"/>
                <rule avp="Origin-Host" required="true" max="1"/>
                <rule avp="Origin-Realm" required="true" max="1"/>
                <rule avp="AVP" required="false"/>
                <rule avp="Failed-AVP" required="false"/>
                <rule avp="Proxy-Info" required="false"/>
                <rule avp="Route-Record" required="false"/>
            </answer>
        </command>

        <command code="318" short="AI" name="Authentication-Information">
            <request>
                <rule avp="Session-Id" required="true" max="1"/>
                <rule avp="Vendor-Specific-Application-Id" required="false" max="1"/>
                <rule avp="Auth-Session-State" required="true" max="1"/>
                <rule avp="Origin-Host" required="true" max="1"/>
                <rule avp="Origin-Realm" required="true" max="1"/>
                <rule avp="Destination-Host" required="false" max="1"/>
                <rule avp="Destination-Realm" required="true" max="1"/>
                <rule avp="User-Name" required="true" max="1"/>
                <rule avp="Supported-Features" required="false"/>
                <rule avp="Requested-EUTRAN-Authentication-Info" required="false" max="1"/>
                <rule avp="Requested-UTRAN-GERAN-Authentication-Info" required="false" max="1"/>
                <rule avp="Visited-PLMN-Id" required="true" max="1"/>
                <rule avp="AVP" required="false"/>
                <rule avp="Proxy-Info" required="false"/>
                <rule avp="Route-Record" required="false"/>
            </request>
            <answer>
                <rule avp="Session-Id" required="true" max="1"/>
                <rule avp="Vendor-Specific-Application-Id" required="false" max="1"/>
                <rule avp="Result-Code" required="false" max="1"/>
                <rule avp="Experimental-Result" required="false" max="1"/>
                <rule avp="Error-Diagnostic" required="false" max="1"/>
                <rule avp="Auth-Session-State" required="true" max="1"/>
                <rule avp="Origin-Host" required="true" max="1"/>
                <rule avp="Origin-Realm" required="true" max="1"/>
                <rule avp="Supported-Features" required="false"/>
                <rule avp="Authentication-Info" required="true" max="1"/>
                <rule avp="AVP" required="false"/>
                <rule avp="Failed-AVP" required="false"/>
                <rule avp="Proxy-Info" required="false"/>
                <rule avp="Route-Record" required="false"/>
            </answer>
        </command>

        <command code="321" short="PU" name="Purge-UE">
            <!--
                < Purge-UE-Request> ::=	< Diameter Header: 321, REQ, PXY, 16777251 >
                < Session-Id >
                [ DRMP ]
                [ Vendor-Specific-Application-Id ]
                { Auth-Session-State }
                { Origin-Host }
                { Origin-Realm }
                [ Destination-Host ]
                { Destination-Realm }
                { User-Name }
                [ OC-Supported-Features ]
                [ PUR-Flags ]
                *[ Supported-Features ]
                [ EPS-Location-Information ]
                *[ AVP ]
                *[ Proxy-Info ]
                *[ Route-Record ]
            -->
            <request>
                <rule avp="Session-Id" required="true" max="1" />
                <rule avp="DRMP" required="false" max="1" />
                <rule avp="Vendor-Specific-Application-Id" required="false" max="1" />
                <rule avp="Auth-Session-State" required="true" max="1" />
                <rule avp="Origin-Host" required="true" max="1" />
                <rule avp="Origin-Realm" required="true" max="1" />
                <rule avp="Destination-Host" required="false" max="1" />
                <rule avp="Destination-Realm" required="true" max="1" />
                <rule avp="User-Name" required="true" max="1" />
                <rule avp="OC-Supported-Features" required="false" max="1" />
                <rule avp="PUR-Flags" required="false" max="1" />
                <rule avp="Supported-Features" required="false" />
                <rule avp="EPS-Location-Information" required="false" max="1" />
                <rule avp="Proxy-Info" required="false" />
                <rule avp="Route-Record" required="false" />
            </request>

            <!--
                < Purge-UE-Answer> ::=	< Diameter Header: 321, PXY, 16777251 >
                < Session-Id >
                [ DRMP ]
                [ Vendor-Specific-Application-Id ]
                *[ Supported-Features ]
                [ Result-Code ]
                [ Experimental-Result ]
                { Auth-Session-State }
                { Origin-Host }
                { Origin-Realm }
                [ OC-Supported-Features ]
                [ OC-OLR ]
                *[ Load ]
                [ PUA-Flags ]
                *[ AVP ]
                [ Failed-AVP ]
                *[ Proxy-Info ]
                *[ Route-Record ]
            -->
            <answer>
                <rule avp="Session-Id" required="true" max="1" />
                <rule avp="DRMP" required="false" max="1" />
                <rule avp="Vendor-Specific-Application-Id" required="false" max="1" />
                <rule avp="Supported-Features" required="false" />
                <rule avp="Result-Code" required="false" max="1" />
                <rule avp="Experimental-Result" required="false" max="1" />
                <rule avp="Auth-Session-State" required="true" max="1" />
                <rule avp="Origin-Host" required="true" max="1" />
                <rule avp="Origin-Realm" required="true" max="1" />
                <rule avp="OC-Supported-Features" required="false" max="1" />
                <rule avp="OC-OLR" required="false" max="1" />
                <!-- rule avp="Load" required="false" /-->
                <rule avp="PUA-Flags" required="false" max="1" />
                <rule avp="Failed-AVP" required="false" max="1" />
                <rule avp="Proxy-Info" required="false" />
                <rule avp="Route-Record" required="false" />
            </answer>
        </command>

        <command code="323" short="NO" name="Notify">
            <!--
                < Notify-Request> ::=	< Diameter Header: 323, REQ, PXY, 16777251 >
                < Session-Id >
                [ Vendor-Specific-Application-Id ]
                [ DRMP ]
                { Auth-Session-State }
                { Origin-Host }
                { Origin-Realm }
                [ Destination-Host ]
                { Destination-Realm }
                { User-Name }
                [ OC-Supported-Features ]
                * [ Supported-Features ]
                [ Terminal-Information ]
                [ MIP6-Agent-Info ]
                [ Visited-Network-Identifier ]
                [ Context-Identifier ]
                [ Service-Selection ]
                [ Alert-Reason ]
                [ UE-SRVCC-Capability ]
                [ NOR-Flags ]
                [ Homogeneous-Support-of-IMS-Voice-Over-PS-Sessions ]
                [ Maximum-UE-Availability-Time ]
                *[ Monitoring-Event-Config-Status ]
                [ Emergency-Services ]
                *[ AVP ]
                *[ Proxy-Info ]
                *[ Route-Record ]
            -->
            <request>
                <rule avp="Session-Id" required="true" max="1" />
                <rule avp="Vendor-Specific-Application-Id" required="false" max="1" />
                <rule avp="DRMP" required="false" max="1" />
                <rule avp="Auth-Session-State" required="true" max="1" />
                <rule avp="Origin-Host" required="true" max="1" />
                <rule avp="Origin-Realm" required="true" max="1" />
                <rule avp="Destination-Host" required="false" max="1" />
                <rule avp="Destination-Realm" required="true" max="1" />
                <rule avp="User-Name" required="true" max="1" />
                <rule avp="OC-Supported-Features" required="false" max="1" />
                <rule avp="Supported-Features" required="false" />
                <rule avp="Terminal-Information" required="false" max="1" />
                <rule avp="MIP6-Agent-Info" required="false" max="1" />
                <rule avp="Visited-Network-Identifier" required="false" max="1" />
                <rule avp="Context-Identifier" required="false" max="1" />
                <rule avp="Service-Selection" required="false" max="1" />
                <rule avp="Alert-Reason" required="false" max="1" />
                <rule avp="UE-SRVCC-Capability" required="false" max="1" />
                <rule avp="NOR-Flags" required="false" max="1" />
                <rule avp="Homogeneous-Support-of-IMS-Voice-Over-PS-Sessions" required="false" max="1" />
                <rule avp="Maximum-UE-Availability-Type" required="false" max="1" />
                <rule avp="Monitoring-Event-Config-Status" required="false" />
                <rule avp="Emergency-Services" required="false" max="1" />
                <rule avp="Proxy-Info" required="false" />
                <rule avp="Route-Record" required="false" />
            </request>
            <!--
                < Notify-Answer> ::=	< Diameter Header: 323, PXY, 16777251 >
                < Session-Id >
                [ DRMP ]
                [ Vendor-Specific-Application-Id ]
                [ Result-Code ]
                [ Experimental-Result ]
                { Auth-Session-State }
                { Origin-Host }
                { Origin-Realm }
                [ OC-Supported-Features ]
                [ OC-OLR ]
                *[ Load ]
                *[ Supported-Features ]
                *[ AVP ]
                [ Failed-AVP ]
                *[ Proxy-Info ]
                *[ Route-Record ]
            -->
            <answer>
                <rule avp="Session-Id" required="true" max="1" />
                <rule avp="DRMP" required="false" max="1" />
                <rule avp="Vendor-Specific-Application-Id" required="false" max="1" />
                <rule avp="Result-Code" required="false" max="1" />
                <rule avp="Experimental-Result" required="false" max="1" />
                <rule avp="Auth-Session-State" required="true" max="1" />
                <rule avp="Origin-Host" required="true" max="1" />
                <rule avp="Origin-Realm" required="true" max="1" />
                <rule avp="OC-Supported-Features" required="false" max="1" />
                <rule avp="OC-OLR" required="false" max="1" />
                <!-- rule avp="Load" required="false" /-->
                <rule avp="Supported-Features" required="false" />
                <rule avp="Failed-AVP" required="false" max="1" />
                <rule avp="Proxy-Info" required="false" />
                <rule avp="Route-Record" required="false" />
            </answer>
        </command>

        <command code="322" short="RS" name="Reset">
            <!--
              < Reset-Request> ::= < Diameter Header: 322, REQ, PXY, 16777251 >
                < Session-Id >
                [ Vendor-Specific-Application-Id ]
                { Auth-Session-State }
                { Origin-Host }
                { Origin-Realm }
                { Destination-Host }
                { Destination-Realm }
                *[ Supported-Features ]
                *[ User-Id ]
                *[ AVP ]
                *[ Proxy-Info ]
                *[ Route-Record ]
            -->
            <request>
                <rule avp="Session-Id" required="true" max="1" />
                <rule avp="Vendor-Specific-Application-Id" required="false" max="1" />
                <rule avp="Auth-Session-State" required="true" max="1" />
                <rule avp="Origin-Host" required="true" max="1" />
                <rule avp="Origin-Realm" required="true" max="1" />
                <rule avp="Destination-Host" required="false" max="1" />
                <rule avp="Destination-Realm" required="true" max="1" />
                <rule avp="Supported-Features" required="false" />
                <rule avp="User-Id" required="false" />
                <rule avp="Proxy-Info" required="false" />
                <rule avp="Route-Record" required="false" />
            </request>
            <!--
              < Reset-Answer> ::= < Diameter Header: 322, PXY, 16777251 >
                < Session-Id >
                [ Vendor-Specific-Application-Id ]
                *[ Supported-Features ]
                [ Result-Code ]
                [ Experimental-Result ]
                { Auth-Session-State }
                { Origin-Host }
                { Origin-Realm }
                *[ AVP ]
                *[ Failed-AVP ]
                *[ Proxy-Info ]
                *[ Route-Record ]
            -->
            <answer>
                <rule avp="Session-Id" required="true" max="1" />
                <rule avp="Vendor-Specific-Application-Id" required="false" max="1" />
                <rule avp="Supported-Features" required="false" />
                <rule avp="Result-Code" required="false" max="1" />
                <rule avp="Experimental-Result" required="false" max="1" />
                <rule avp="Auth-Session-State" required="true" max="1" />
                <rule avp="Origin-Host" required="true" max="1" />
                <rule avp="Origin-Realm" required="true" max="1" />
                <rule avp="Failed-AVP" required="false" max="1" />
                <rule avp="Proxy-Info" required="false" />
                <rule avp="Route-Record" required="false" />
            </answer>
        </command>

        <avp name="Supported-Features" code="628" vendor-id="10415" must="V" may="M" may-encrypt="N">
            <data type="Grouped">
                <rule avp="Vendor-Id" required="true" max="1"/>
                <rule avp="Feature-List-ID" required="true" max="1"/>
                <rule avp="Feature-List" required="true" max="1"/>
            </data>
        </avp>

        <avp name="Subscription-Data" code="1400" vendor-id="10415" must="M,V" may-encrypt="N">
            <data type="Grouped">
                <rule avp="Subscriber-Status" required="false" max="1"/>
                <rule avp="MSISDN" required="false" max="1"/>
                <rule avp="STN-SR" required="false" max="1"/>
                <rule avp="ICS-Indicator" required="false" max="1"/>
                <rule avp="Network-Access-Mode" required="false" max="1"/>
                <rule avp="Operator-Determined-Barring" required="false" max="1"/>
                <rule avp="HPLMN-ODB" required="false" max="1"/>
                <rule avp="Regional-Subscription-Zone-Code" required="false" max="10"/>
                <rule avp="Access-Restriction-Data" required="false" max="1"/>
                <rule avp="APN-OI-Replacement" required="false" max="1"/>
                <rule avp="LCS-Info" required="false" max="1"/>
                <rule avp="Teleservice-List" required="false" max="1"/>
                <rule avp="Call-Barring-Info" required="false"/>
                <rule avp="TGPP-Charging-Characteristics" required="false" max="1"/>
                <rule avp="AMBR" required="false" max="1"/>
                <rule avp="APN-Configuration-Profile" required="false" max="1"/>
                <rule avp="RAT-Frequency-Selection-Priority-ID" required="false" max="1"/>
                <rule avp="Trace-Data" required="false" max="1"/>
                <rule avp="GPRS-Subscription-Data" required="false" max="1"/>
                <rule avp="CSG-Subscription-Data" required="false"/>
                <rule avp="Roaming-Restricted-Due-To-Unsupported-Feature" required="false" max="1"/>
                <rule avp="Subscribed-Periodic-RAU-TAU-Timer" required="false" max="1"/>
                <rule avp="MPS-Priority" required="false" max="1"/>
                <rule avp="VPLMN-LIPA-Allowed" required="false" max="1"/>
                <rule avp="Relay-Node-Indicator" required="false" max="1"/>
                <rule avp="MDT-User-Consent" required="false" max="1"/>
                <rule avp="AVP" required="false"/>
            </data>
        </avp>

        <avp name="Subscriber-Status" code="1424" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="SERVICE_GRANTED"/>
                <item code="1" name="OPERATOR_DETERMINED_BARRING"/>
            </data>
        </avp>

        <avp name="STN-SR" code="1433" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="OctetString"/>
        </avp>

        <avp name="ICS-Indicator" code="1491" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="FALSE"/>
                <item code="1" name="TRUE"/>
            </data>
        </avp>

        <avp name="Network-Access-Mode" code="1417" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="PACKET_AND_CIRCUIT"/>
                <item code="1" name="RESERVED"/>
                <item code="2" name="ONLY_PACKET"/>
            </data>
        </avp>

        <avp name="Operator-Determined-Barring" code="1425" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Unsigned32"/>
        </avp>

        <avp name="HPLMN-ODB" code="1418" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Unsigned32"/>
        </avp>

        <avp name="Regional-Subscription-Zone-Code" code="1446" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="OctetString"/>
        </avp>

        <avp name="Access-Restriction-Data" code="1426" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Unsigned32"/>
        </avp>

        <avp name="APN-OI-Replacement" code="1427" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="UTF8String"/>
        </avp>

        <avp name="LCS-Info" code="1473" vendor-id="10415" must="M,V" may-encrypt="N">
            <data type="Grouped">
                <rule avp="GMLC-Number" required="false"/>
                <rule avp="LCS-PrivacyException" required="false"/>
                <rule avp="MO-LR" required="false"/>
                <rule avp="AVP" required="false"/>
            </data>
        </avp>avp>

        <avp name="GMLC-Number" code="1474" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="OctetString"/>
        </avp>

        <avp name="LCS-PrivacyException" code="1475" vendor-id="10415" must="M,V" may-encrypt="N">
            <data type="Grouped">
                <rule avp="SS-Code" required="true" max="1"/>
                <rule avp="SS-Status" required="true" max="1"/>
                <rule avp="Notification-To-UE-User" required="false" max="1"/>
                <rule avp="External-Client" required="false"/>
                <rule avp="PLMN-Client" required="false"/>
                <rule avp="TGPP-Service-Type" required="false"/>
                <rule avp="AVP" required="false"/>
            </data>
        </avp>avp>

        <avp name="SS-Code" code="1476" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="OctetString"/>
        </avp>

        <avp name="SS-Status" code="1477" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="OctetString"/>
        </avp>

        <avp name="Notification-To-UE-User" code="1478" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="NOTIFY_LOCATION_ALLOWED"/>
                <item code="1" name="NOTIFYANDVERIFY_LOCATION_ALLOWED_IF_NO_RESPONSE"/>
                <item code="2" name="NOTIFYANDVERIFY_LOCATION_NOT_ALLOWED_IF_NO_RESPONSE"/>
                <item code="3" name="LOCATION_NOT_ALLOWED"/>
            </data>
        </avp>

        <avp name="External-Client" code="1479" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Grouped">
                <rule avp="Client-Identity" required="true" max="1"/>
                <rule avp="GMLC-Restriction" required="false" max="1"/>
                <rule avp="Notification-To-UE-User" required="false" max="1"/>
                <rule avp="AVP" required="false"/>
            </data>
        </avp>avp>

        <avp name="Client-Identity" code="1480" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="OctetString"/>
        </avp>

        <avp name="GMLC-Restriction" code="1481" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="GMLC_LIST"/>
                <item code="1" name="HOME_COUNTRY"/>
            </data>
        </avp>

        <avp name="PLMN-Client" code="1482" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="BROADCAST_SERVICE"/>
                <item code="1" name="O_AND_M_HPLMN"/>
                <item code="2" name="O_AND_M_VPLMN"/>
                <item code="3" name="ANONYMOUS_LOCATION"/>
                <item code="3" name="TARGET_UE_SUBSCRIBED_SERVICE"/>
            </data>
        </avp>

        <avp name="TGPP-Service-Type" code="1483" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Grouped">
                <rule avp="ServiceTypeIdentity" required="true" max="1"/>
                <rule avp="GMLC-Restriction" required="false" max="1"/>
                <rule avp="Notification-To-UE-User" required="false" max="1"/>
                <rule avp="AVP" required="false"/>
            </data>
        </avp>

        <avp name="ServiceTypeIdentity" code="1484" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Unsigned32"/>
        </avp>

        <avp name="MO-LR" code="1485" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Grouped">
                <rule avp="SS-Code" required="true" max="1"/>
                <rule avp="SS-Status" required="true" max="1"/>
                <rule avp="AVP" required="false"/>
            </data>
        </avp>

        <avp name="Teleservice-List" code="1486" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Grouped">
                <rule avp="TS-Code" required="true" max="1"/>
                <rule avp="AVP" required="false"/>
            </data>
        </avp>

        <avp name="TS-Code" code="1487" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="OctetString"/>
        </avp>

        <avp name="Call-Barring-Info" code="1488" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Grouped">
                <rule avp="SS-Code" required="true" max="1"/>
                <rule avp="SS-Status" required="true" max="1"/>
                <rule avp="AVP" required="false"/>
            </data>
        </avp>

        <avp name="AMBR" code="1435" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Grouped">
                <rule avp="Max-Requested-Bandwidth-UL" required="true" max="1"/>
                <rule avp="Max-Requested-Bandwidth-DL" required="true" max="1"/>
                <rule avp="AVP" required="false"/>
            </data>
        </avp>

        <avp name="APN-Configuration-Profile" code="1429" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Grouped">
                <rule avp="Context-Identifier" required="true" max="1"/>
                <rule avp="All-APN-Configurations-Included-Indicator" required="true" max="1"/>
                <rule avp="APN-Configuration" required="true"/>
                <rule avp="AVP" required="false"/>
            </data>
        </avp>

        <avp name="Context-Identifier" code="1423" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Unsigned32"/>
        </avp>

        <avp name="All-APN-Configurations-Included-Indicator" code="1428" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="All_APN_CONFIGURATIONS_INCLUDED"/>
                <item code="1" name="MODIFIED|ADDED_APN_CONFIGURATIONS_INCLUDED"/>
            </data>
        </avp>

        <avp name="APN-Configuration" code="1430" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Grouped">
                <rule avp="Context-Identifier" required="true" max="1"/>
                <rule avp="Served-Party-IP-Address" required="false" max="2"/>
                <rule avp="PDN-Type" required="true" max="1"/>
                <rule avp="Service-Selection" required="true" max="1"/>
                <rule avp="EPS-Subscribed-QoS-Profile" required="false" max="1"/>
                <rule avp="VPLMN-Dynamic-Address-Allowed" required="false" max="1"/>
                <rule avp="MIP6-Agent-Info" required="false" max="1"/>
                <rule avp="Visited-Network-Identifier" required="false" max="1"/>
                <rule avp="PDN-GW-Allocation-Type" required="false" max="1"/>
                <rule avp="TGPP-Charging-Characteristics" required="false" max="1"/>
                <rule avp="AMBR" required="false" max="1"/>
                <rule avp="Specific-APN-Info" required="false"/>
                <rule avp="APN-OI-Replacement" required="false" max="1"/>
                <rule avp="SIPTO-Permission" required="false" max="1"/>
                <rule avp="LIPA-Permission" required="false" max="1"/>
                <rule avp="AVP" required="false"/>
            </data>
        </avp>

        <avp name="Served-Party-IP-Address" code="848" must="M,V" may="P" may-encrypt="N" vendor-id="10415">
            <data type="Address"/>
        </avp>

        <avp name="PDN-Type" code="1456" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="IPv4"/>
                <item code="1" name="IPv6"/>
                <item code="2" name="IPv4v6"/>
                <item code="3" name="IPv4_OR_IPv6"/>
            </data>
        </avp>

        <avp name="EPS-Subscribed-QoS-Profile" code="1431" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Grouped">
                <rule avp="QoS-Class-Identifier" required="true" max="1"/>
                <rule avp="Allocation-Retention-Priority" required="true" max="1"/>
                <rule avp="AVP" required="false"/>
            </data>
        </avp>

        <avp name="VPLMN-Dynamic-Address-Allowed" code="1432" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="NOTALLOWED"/>
                <item code="1" name="ALLOWED"/>
            </data>
        </avp>

        <avp name="PDN-GW-Allocation-Type" code="1438" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="STATIC"/>
                <item code="1" name="DYNAMIC"/>
            </data>
        </avp>

        <avp name="SIPTO-Permission" code="1613" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="SIPTO_ALLOWED"/>
                <item code="1" name="SIPTO_NOTALLOWED"/>
            </data>
        </avp>

        <avp name="LIPA-Permission" code="1618" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="LIPA-PROHIBITED"/>
                <item code="1" name="LIPA-ONLY"/>
                <item code="2" name="LIPA-CONDITIONAL"/>
            </data>
        </avp>

        <avp name="RAT-Frequency-Selection-Priority-ID" code="1440" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Unsigned32"/>
        </avp>

        <avp name="Trace-Data" code="1458" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Grouped">
                <rule avp="Trace-Reference" required="true" max="1"/>
                <rule avp="Trace-Depth" required="true" max="1"/>
                <rule avp="Trace-NE-Type-List" required="true" max="1"/>
                <rule avp="Trace-Interface-List" required="false" max="1"/>
                <rule avp="Trace-Event-List" required="true" max="1"/>
                <rule avp="OMC-Id" required="false" max="1"/>
                <rule avp="Trace-Collection-Entity" required="true" max="1"/>
                <rule avp="MDT-Configuration" required="false" max="1"/>
                <rule avp="AVP" required="false"/>
            </data>
        </avp>avp>

        <avp name="Trace-Reference" code="1459" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="OctetString"/>
        </avp>

        <avp name="Trace-Depth" code="1462" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="LIPA-PROHIBITED"/>
                <item code="1" name="LIPA-ONLY"/>
                <item code="2" name="LIPA-CONDITIONAL"/>
            </data>
        </avp>

        <avp name="Trace-NE-Type-List" code="1463" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="OctetString"/>
        </avp>

        <avp name="Trace-Interface-List" code="1464" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="OctetString"/>
        </avp>

        <avp name="Trace-Event-List" code="1465" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="OctetString"/>
        </avp>

        <avp name="OMC-Id" code="1466" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="OctetString"/>
        </avp>

        <avp name="Trace-Event-List" code="1465" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="OctetString"/>
        </avp>

        <avp name="Trace-Collection-Entity" code="1452" must="M,V" may="P" may-encrypt="N" vendor-id="10415">
            <data type="Address"/>
        </avp>

        <avp name="MDT-Configuration" code="1622" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Grouped">
                <rule avp="QoS-Class-Identifier" required="true" max="1"/>
                <rule avp="Allocation-Retention-Priority" required="true" max="1"/>
                <rule avp="AVP" required="false"/>
            </data>
        </avp>

        <avp name="GPRS-Subscription-Data" code="1467" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Grouped">
                <rule avp="Complete-Data-List-Included-Indicator" required="true" max="1"/>
                <rule avp="PDP-Context" required="true" max="50"/>
                <rule avp="AVP" required="false"/>
            </data>
        </avp>

        <avp name="Complete-Data-List-Included-Indicator" code="1468" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="All_PDP_CONTEXTS_INCLUDED"/>
                <item code="1" name="MODIFIED/ADDED_PDP CONTEXTS_INCLUDED"/>
            </data>
        </avp>

        <avp name="PDP-Context" code="1469" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Grouped">
                <rule avp="Context-Identifier" required="true" max="1"/>
                <rule avp="PDP-Type" required="true" max="1"/>
                <rule avp="PDP-Address" required="false" max="1"/>
                <rule avp="QoS-Subscribed" required="true" max="1"/>
                <rule avp="VPLMN-Dynamic-Address-Allowed" required="false" max="1"/>
                <rule avp="Service-Selection" required="true" max="1"/>
                <rule avp="TGPP-Charging-Characteristics" required="false" max="1"/>
                <rule avp="Ext-PDP-Type" required="false" max="1"/>
                <rule avp="Ext-PDP-Address" required="false" max="1"/>
                <rule avp="AMBR" required="false" max="1"/>
                <rule avp="SIPTO-Permission" required="false" max="1"/>
                <rule avp="LIPA-Permission" required="false" max="1"/>
                <rule avp="AVP" required="false"/>
            </data>
        </avp>

        <avp name="PDP-Type" code="1470" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="OctetString"/>
        </avp>

        <avp name="QoS-Subscribed" code="1404" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="OctetString"/>
        </avp>

        <avp name="VPLMN-Dynamic-Address-Allowed" code="1432" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="NOTALLOWED"/>
                <item code="1" name="ALLOWED"/>
            </data>
        </avp>

        <avp name="Ext-PDP-Type" code="1620" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="OctetString"/>
        </avp>

        <avp name="Ext-PDP-Address" code="1621" must="V,M" may="P" must-not="-" may-encrypt="Y" vendor-id="10415">
            <data type="Address"/>
        </avp>

        <avp name="SIPTO-Permission" code="1613" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="SIPTO_ALLOWED"/>
                <item code="1" name="SIPTO_NOTALLOWED"/>
            </data>
        </avp>

        <avp name="CSG-Subscription-Data" code="1436" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Grouped">
                <rule avp="CSG-Id" required="true" max="1"/>
                <rule avp="Expiration-Date" required="false" max="1"/>
                <rule avp="Service-Selection" required="false"/>
                <rule avp="AVP" required="false"/>
            </data>
        </avp>

        <avp name="Expiration-Date" code="1439" must="V,M" may-encrypt="N" vendor-id="10415">
            <data type="Time"/>
        </avp>

        <avp name="Roaming-Restricted-Due-To-Unsupported-Feature" code="1457" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="Roaming-Restricted-Due-To-Unsupported-Feature"/>
            </data>
        </avp>

        <avp name="Subscribed-Periodic-RAU-TAU-Timer" code="1619" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="Unsigned32"/>
        </avp>

        <avp name="MPS-Priority" code="1616" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="Unsigned32"/>
        </avp>

        <avp name="VPLMN-LIPA-Allowed" code="1617" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="LIPA-NOTALLOWED"/>
                <item code="1" name="LIPA-ALLOWED"/>
            </data>
        </avp>

        <avp name="Relay-Node-Indicator" code="1633" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="NOT_RELAY_NODE"/>
                <item code="1" name="RELAY_NODE"/>
            </data>
        </avp>

        <avp name="MDT-User-Consent" code="1634" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="CONSENT_NOT_GIVEN"/>
                <item code="1" name="CONSENT_GIVEN"/>
            </data>
        </avp>

        <avp name="Feature-List-ID" code="629" must="V" must_not="M" may-encrypt="N" vendor-id="10415">
            <data type="Unsigned32"/>
        </avp>

        <avp name="Feature-List" code="630" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="Unsigned32"/>
        </avp>

        <avp name="Requested-EUTRAN-Authentication-Info" code="1408" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Grouped">
                <rule avp="Number-Of-Requested-Vectors" required="false" max="1"/>
                <rule avp="Immediate-Response-Preferred" required="false" max="1"/>
                <rule avp="Re-synchronization-Info" required="false" max="1"/>
            </data>
        </avp>

        <avp name="Number-Of-Requested-Vectors" code="1410" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Unsigned32"/>
        </avp>

        <avp name="Re-synchronization-Info" code="1411" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="OctetString"/>
        </avp>

        <avp name="Immediate-Response-Preferred" code="1412" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Unsigned32"/>
        </avp>

        <avp name="Requested-UTRAN-GERAN-Authentication-Info" code="1409" vendor-id="10415" must="M,V" may-encrypt="N">
            <data type="Grouped">
                <rule avp="Number-Of-Requested-Vectors" required="false" max="1"/>
                <rule avp="Immediate-Response-Preferred" required="false" max="1"/>
                <rule avp="Re-synchronization-Info" required="false" max="1"/>
            </data>
        </avp>

        <avp name="Visited-PLMN-Id" code="1407" must="MV" may-encrypt="N" vendor-id="10415">
            <data type="OctetString"/>
        </avp>

        <avp name="Error-Diagnostic" code="1614" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="GPRS_DATA_SUBSCRIBED"/>
                <item code="1" name="NO_GPRS_DATA_SUBSCRIBED"/>
                <item code="2" name="ODB-ALL-APN"/>
                <item code="3" name="ODB-HPLMN-APN"/>
                <item code="4" name="ODB-VPLMN-APN"/>
            </data>
        </avp>

        <avp name="Authentication-Info" code="1413" vendor-id="10415" must="M,V" may-encrypt="N">
            <data type="Grouped">
                <rule avp="E-UTRAN-Vector" required="false"/>
                <rule avp="UTRAN-Vector" required="false"/>
                <rule avp="GERAN-Vector" required="false"/>
            </data>
        </avp>

        <avp name="E-UTRAN-Vector" code="1414" vendor-id="10415" must="M.V" may-encrypt="N">
            <data type="Grouped">
                <rule avp="Item-Number" required="false" max="1"/>
                <rule avp="RAND" required="true" max="1"/>
                <rule avp="XRES" required="true" max="1"/>
                <rule avp="AUTN" required="true" max="1"/>
                <rule avp="KASME" required="true" max="1"/>
            </data>
        </avp>

        <avp name="UTRAN-Vector" code="1415" vendor-id="10415" must="M,V" may-encrypt="N">
            <data type="Grouped">
                <rule avp="Item-Number" required="false"/>
                <rule avp="RAND" required="true"/>
                <rule avp="XRES" required="true"/>
                <rule avp="AUTN" required="true"/>
                <rule avp="Confidentiality-Key" required="true"/>
                <rule avp="Integrity-Key" required="true"/>
            </data>
        </avp>

        <avp name="GERAN-Vector" code="1416" vendor-id="10415" must="M,V" may-encrypt="N">
            <data type="Grouped">
                <rule avp="Item-Number" required="false" max="1"/>
                <rule avp="RAND" required="true" max="1"/>
                <rule avp="XRES" required="true" max="1"/>
                <rule avp="Kc" required="true" max="1"/>
            </data>
        </avp>

        <avp name="Item-Number" code="1419" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Unsigned32"/>
        </avp>

        <avp name="Cancellation-Type" code="1420" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="MME_UPDATE_PROCEDURE"/>
                <item code="1" name="SGSN_UDPATE_PROCEDURE"/>
                <item code="2" name="SUBSCRIPTION_WITHDRAWAL"/>
                <item code="3" name="UPDATE_PROCEDURE_IWF"/>
                <item code="4" name="INITIAL_ATTACH_PROCEDURE"/>
            </data>
        </avp>

        <avp name="RAND" code="1447" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="OctetString"/>
        </avp>

        <avp name="XRES" code="1448" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="OctetString"/>
        </avp>

        <avp name="AUTN" code="1449" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="OctetString"/>
        </avp>

        <avp name="KASME" code="1450" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="OctetString"/>
        </avp>

        <avp name="Kc" code="1453" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="OctetString"/>
        </avp>

        <avp name="Confidentiality-Key" code="625" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="OctetString"/>
        </avp>

        <avp name="Integrity-Key" code="626" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="OctetString"/>
        </avp>

        <avp name="ULR-Flags" code="1405" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Unsigned32"/>
        </avp>

        <avp name="ULA-Flags" code="1406" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Unsigned32"/>
        </avp>

        <avp name="CLR-Flags" code="1638" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="Unsigned32"/>
        </avp>

        <avp name="UE-SRVCC-Capability" code="1615" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="UE-SRVCC-NOT-SUPPORTED"/>
                <item code="1" name="UE-SRVCC-SUPPORTED"/>
            </data>
        </avp>

        <avp name="Homogeneous-Support-of-IMS-Voice-Over-PS-Sessions" code="1493" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="NOT-SUPPORTED"/>
                <item code="1" name="SUPPORTED"/>
            </data>
        </avp>

        <avp name="Active-APN" code="1612" vendor-id="10415" must="M,V" may-encrypt="N">
            <data type="Grouped">
                <rule avp="Context-Identifier" required="true" max="1"/>
                <rule avp="Service-Selection" required="false" max="1"/>
                <rule avp="MIP6-Agent-Info" required="false" max="1"/>
                <rule avp="Visited-Network-Identifier" required="false" max="1"/>
                <rule avp="Specific-APN-Info" required="false" max="1"/>
                <rule avp="AVP" required="false"/>
            </data>
        </avp>

        <avp name="Specific-APN-Info" code="1472" vendor-id="10415" must="M,V" may-encrypt="N">
            <data type="Grouped">
                <rule avp="Service-Selection" required="true" max="1"/>
                <rule avp="MIP6-Agent-Info" required="true" max="1"/>
                <rule avp="Visited-Network-Identifier" required="false" max="1"/>
                <rule avp="AVP" required="false"/>
            </data>
        </avp>

        <avp name="Context-Identifier" code="1423" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Unsigned32"/>
        </avp>

        <avp name="PUR-Flags" code="1635" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="Unsigned32" />
        </avp>

        <avp name="PUA-Flags" code="1442" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Unsigned32" />
        </avp>

        <avp name="NOR-Flags" code="1443" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Unsigned32" />
        </avp>

        <avp name="Subscribed-VSRVCC" code="1636" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="VSRVCC_SUBSCRIBED" />
            </data>
        </avp>

        <avp name="MIP-Home-Agent-Address" code="334" must="M" must-not="V" vendor-id="10415">
            <data type="Address"/>
        </avp>

        <!-- RFC 4004 -->
        <avp name="MIP-Home-Agent-Host" code="348" must="M" may="P" must-not="V" may-encrypt="Y" vendor-id="10415">
            <data type="Grouped">
                <rule avp="Destination-Realm" required="true" max="1"/>
                <rule avp="Destination-Host" required="true" max="1"/>
                <rule avp="AVP" required="false"/>
            </data>
        </avp>

        <!-- RFC 5447 Diameter Mobile IPv6: Support for Network Access Server to Diameter Server Interaction -->
        <avp name="MIP6-Agent-Info" code="486" must="M" may="P" must-not="V" may-encrypt="Y" vendor-id="10415">
            <data type="Grouped">
                <rule avp="MIP-Home-Agent-Address" required="false" max="2"/>
                <rule avp="MIP-Home-Agent-Host" required="false" max="1"/>
                <rule avp="MIP6-Home-Link-Prefix" required="false" max="1"/>
                <rule avp="AVP" required="false"/>
            </data>
        </avp>

        <avp name="Service-Selection" code="493" must="M" may="P" must-not="V" may-encrypt="Y" vendor-id="10415">
            <data type="UTF8String"/>
        </avp>

        <avp name="Visited-Network-Identifier" code="600" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="OctetString"/>
        </avp>

        <avp name="MIP6-Home-Link-Prefix" code="125" vendor-id="10415">
            <data type="OctetString"/>
        </avp>

        <avp name="User-Id" code="1444" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="UTF8String"/>
        </avp>

    </application>
</diameter>`
