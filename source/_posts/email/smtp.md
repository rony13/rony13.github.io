---
title: Simple Mail Transfer Protocol (SMTP)
categories: email
date: 2019-11-28 14:48:16
---

## Overview 
SMTP is a __connection-oriented__, __text-based protocol__ in which a mail sender communicates with a mail receiver by issuing command strings and supplying necessary data over a reliable ordered data stream channel, typically a TCP connection.

SMTP defines __message transport__, not the message content. Thus, it defines the mail envelope and its parameters, such as the envelope sender, but not the header (except trace information) nor the body of the message itself.

SMTP is a __Application Layer Protocal__. E-mail is submitted by a __MUA__(mail user agent) to a __MSA__(mail submission agent) using __SMTP__ on __TCP__. The MSA delivers the mail to its __MTA__(MTA). Often, these two agents are instances of the same software launched with different options on the same machine. Local processing can be done either on a single machine, or split among multiple machines; mail agent processes on one machine can share files, but if processing is on multiple machines, they transfer messages between each other using SMTP, where each machine is configured to use the next machine as a smart host. Each process is an __MTA__(an SMTP server) in its own right.

The boundary MTA uses the __DNS__(Domain name system) to look up the __MX Record__(mail exchanger record) for the recipient's domain (the part of the email address on the right of @). The MX record contains the name of the target host. Based on the target host and other factors, the MTA selects an exchange server: see the article MX record. The MTA connects to the exchange server as an SMTP client.

Message transfer can occur in a single connection between two MTAs, or in a series of hops through intermediary systems. A receiving SMTP server may be the ultimate destination, an intermediate "relay" (that is, it stores and forwards the message) or a "gateway" (that is, it may forward the message using some protocol other than SMTP). Each hop is a formal handoff of responsibility for the message, whereby the receiving server must either deliver the message or properly report the failure to do so.

Once the final hop accepts the incoming message, it hands it to a __MDA__(mail delivery agent) for local delivery. An MDA saves messages in the relevant mailbox format. As with sending, this reception can be done using one or multiple computers, but in the diagram above the MDA is depicted as one box near the mail exchanger box. An MDA may deliver messages directly to storage, or forward them over a network using SMTP or other protocol such as __LMTP__(Local Mail Transfer Protocol), a derivative of SMTP designed for this purpose.

Once delivered to the local mail server, the mail is stored for batch retrieval by authenticated mail clients (MUAs). Mail is retrieved by end-user applications, called email clients, using __IMAP__(Internet Message Access Protocol), a protocol that both facilitates access to mail and manages stored mail, or the __POP3__(Post Office Protocol Version 3) which typically uses the traditional mbox mail file format or a proprietary system such as Microsoft Exchange/Outlook or Lotus Notes/Domino. Webmail clients may use either method, but the retrieval protocol is often not a formal standard.

![overview](/images/smtp-overview.png)

## MX Record

A MX record(mail exchanger record) specifies the mail server responsible for accepting email messages on behalf of a domain name. It is a resource record in the DNS(Domain Name System). It is possible to configure several MX records, typically pointing to an array of mail servers for load balancing and redundancy.

You can search for the MX records for a specific domain with __nslookup__ on mac:
```
# nslookup
> set q=mx
> gmail.com
Server:		10.86.96.1
Address:	10.86.96.1#53

Non-authoritative answer:
gmail.com	mail exchanger = 40 alt4.gmail-smtp-in.l.google.com.
gmail.com	mail exchanger = 5 gmail-smtp-in.l.google.com.
gmail.com	mail exchanger = 10 alt1.gmail-smtp-in.l.google.com.
gmail.com	mail exchanger = 20 alt2.gmail-smtp-in.l.google.com.
gmail.com	mail exchanger = 30 alt3.gmail-smtp-in.l.google.com.

Authoritative answers can be found from:

```
## Authentication
Authencication, Authorization and Audit is a big topic, we'll discuss later.

## Port

There're three ports for SMTP
1. Port 25
the oldest port. It's suggested to be used between MTAs in order to against spam E-mail.

2. Port 587
You need authentication when use port 587, that's why it's recomanded to use between MUA and MSA.

3. Port 465
Deprecated.

![port](/images/smtp-port.png)


## Transport Example
```
~ telnet smtp.163.com 25
Trying 220.181.12.12...
Connected to smtp.163.com.
Escape character is '^]'.
220 163.com Anti-spam GT for Coremail System (163com[20141201])
EHLO SMTP
250-mail
250-PIPELINING
250-AUTH LOGIN PLAIN
250-AUTH=LOGIN PLAIN
250-coremail 1Uxr2xKj7kG0xkI17xGrU7I0s8FY2U3Uj8Cz28x1UUUUU7Ic2I0Y2UFJbpo8UCa0xDrUUUUj
250-STARTTLS
250 8BITMIME
AUTH LOGIN
334 dXNlcm5hbWU6
dG9ueUAxNjMuY29t
334 UGFzc3dvcmQ6
dGVzdA==
535 Error: authentication failed
```

