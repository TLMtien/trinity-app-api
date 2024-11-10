
![schema-database](/public/schema-database.png)

### 1. Local Setup Guide - Build project on local (using docker):
*	Create docker network
```sh
	$ make create_network
```
*	Build docker
```sh
	$ make docker_build
```
*	migrations postgres database
```sh
	$ make migrate_up
```
### 2. API document
*  Access swagger link:
	http://127.0.0.1:8080/swagger/index.html

* Technologies: Gin, gorm, postgres, docker

### 3. Module
1. Campaign module: Module manages marketing campaigns and checks their eligibility before creating any vouchers.
  - Attributes:
    - Name: The name of the campaign.
    - Start Date: The date when the campaign becomes active.
    -	End Date: The date when the campaign ends.
    -	Max Vouchers: The total number of vouchers that can be generated for this campaign.
    -	Available Vouchers: The current number of vouchers that can still be generated (initially equal to Max Vouchers).
    -	Holding Vouchers: The number of vouchers that have been issued but not yet redeemed (initially 0).
  - Actions:
    -	When creating a voucher, the system checks if the campaign is eligible.
    -	If the voucher is created successfully:
    -	Increase the Holding Vouchers count by 1.
    -	Decrease the Available Vouchers count by 1.

2. Voucher Module: module handles the creation and management of vouchers tied to  campaigns
  - Attributes:
    -	Code: A unique code generated for the campaign.
    -	User ID: The ID of the user who redeems the voucher. If User ID is null, it means the voucher has not yet been used. A voucher is considered "holding" if its User ID is null and it is still active.
    -	Campaign ID: The campaign to which this voucher belongs.
    -	Discount Percentage: The maximum discount percentage provided by this voucher.
    -	Expired At: The expiration date of the voucher.
    -	Is Active: Indicates whether the voucher is still valid. A voucher can become inactive immediately if its campaign is no longer eligible, even before its expiration date.
	
  - Actions:
    -	When creating a voucher, verify that the campaign is eligible.
    -	If the voucher is created successfully:
    -	Increment the Holding Vouchers count for the campaign.
    -	Decrement the Available Vouchers count for the campaign.

3. Purchase module: module manages user purchases
  - Attributes:
    -	User ID: The ID of the user making the purchase.
    -	Voucher Code: The voucher code being used in the purchase.
    -	Subscription Plan Price Details ID: The ID of the selected subscription plan.
    -	Total Price: The final price after applying any discounts.
  - Actions:
    -	When creating a purchase, validate the voucher.
    -	If the voucher is valid and the purchase is successful:
    -	Update the User ID on the voucher (indicating the voucher has been redeemed).
    -	Decrease the Holding Vouchers count for the associated campaign (which also indirectly reduces the Available Vouchers count).

### 4. Improvements
1. Voucher user limit: each voucher can only be redeemed by a single user. Once a voucher is used, it cannot be assigned to another user.
2. Cron job for expired unused vouchers: a scheduled cron job will periodically check for vouchers that have expired but have not yet been redeemed. 
   - For each expired unused voucher:
     - The Holding Vouchers count will be decreased (since the voucher is no longer valid).
     - The Available Vouchers count will be increased, making the voucher available for new issuance.

