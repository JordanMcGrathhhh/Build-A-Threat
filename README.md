# Build-A-Threat

**Build-A-Threat** is a Golang application designed to enhance your threat detection and response skills using the MITRE ATT&CK® framework. This tool takes data from MITRE's ATT&CK framework via TAXII (Trusted Automated eXchange of Indicator Information), randomly selects three techniques, and challenges you to analyze them in the context of the attack chain. 

The primary goal of Build-A-Threat is to help cybersecurity professionals, SOC analysts, and Threat Detection Engineers think critically about detection strategies and improve their ability to recognize malicious behaviors.

---

## Features

- **TAXII Integration**: Fetches the latest techniques and tactics directly from the MITRE ATT&CK framework.
- **Random Technique Selection**: Selects three techniques at random to simulate a realistic attack scenario.
- **Interactive Challenge**: Encourages users to map the techniques to the attack chain, infer attacker goals, and devise detection strategies.
- **Educational Tool**: A practice tool for sharpening your understanding of attack patterns and threat detection methodologies without any pressure.

---

## How It Works

1. **Fetch Data**: The application connects to MITRE ATT&CK’s TAXII server to retrieve up-to-date techniques.
2. **Random Selection**: Three techniques are chosen at random from the available data.
3. **Individual Exercise**: Think critically about the techniques, map them to the ATT&CK phase, and brainstorm detection methods or mitigation strategies.
5. **Reflection**: Compare your thought process to documented mitigations and detections for the selected techniques.
