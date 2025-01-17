class Solution:
    def numUniqueEmails(self, emails: list[str]) -> int:
        unique_emails = set()
        for i in range(len(emails)):
            cleared_email = []
            at = False
            plus = False
            for symbol in emails[i]:
                if symbol == "@":
                    at = True
                    plus = False
                elif not at:
                    if symbol == "+":
                        plus = True
                    elif symbol == ".":
                        continue

                if at or not plus:
                    cleared_email.append(symbol)
            email = "".join(cleared_email)
            unique_emails.add(email)
        return len(unique_emails)
