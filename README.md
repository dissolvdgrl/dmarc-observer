# DMARC Report Viewer

A silly little command line program that scans a directory on our computer for DMARC reports and prints out a basic summary. You can have DMARC reports sent to your email address if you configure it in your DNS.

I am still building this out, not done yet!

## What is DMARC?
Domain-based Message Authentication, Reporting, and Conformance. They are feedback messages sent by email service providers like Yahoo, Google or Outlook that shows how emails sent from your domain are being processed by these providers.

It shows who is sending emails on your behalf, whether they pass security checks, and if anyone is trying to spoof or fake your domain. This can help you fix delivery problems with email

## Why did I build this?
Out of curiosity, and I wanted to build something with Go to keep my skills sharp, and maybe I think DMARC reports are kinda interesting for some reason. Fuck LLMs, btw.


![CLI interface]([image-url-or-path](https://chilldsgn.com/assets/screenshot-from-2026-10-09-16-04-53.png))
