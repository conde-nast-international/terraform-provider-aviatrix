# Acceptance Tests

#### Pre-requisites

- The controller must be launched before hand and must be up and running the latest controller version
- IAM roles (aviatrix-role-ec2 and aviatrix-role-app) also must be created and attached if any IAM role related tests are to be run. Currently all tests are based on Access key, Secret key
- The VPC's with public subnet to launch the gateways must be created before the tests.
- If you are running aviatrix_aws_peer or aviatrix_peer, two VPC's with non overlapping CIDR's must be created before hand
- If you are running the tests on a BYOL controller, the customer ID must be set prior to the tests, otherwise run the tests on a PayG metered controller.
- aviatrix_dc_extn test can only be run on CloudN platform and not on UCC. CloudN must also be initialized and subnets must be configured before the tests are run

#### Skip parameters and variables

Passing an environment value of "yes" to the skip parameter allows you to skip the particular resource. If it is not skipped, it checks for the existence of other required variables. Generic variables are required for any acceptance test

| Test module name      | Skip parameter    | Required variables                                           |
| --------------------- | ----------------- | ------------------------------------------------------------ |
| Generic               | N/A               | AVIATRIX_USERNAME, AVIATRIX_PASSWORD, AVIATRIX_CONTROLLER_IP |
| aviatrix_account      | SKIP_ACCOUNT      | AWS_ACCOUNT_NUMBER, AWS_ACCESS_KEY, AWS_SECRET_KEY           |
| aviatrix_account_user | SKIP_ACCOUNT_USER |                                                              |
| aviatrix_aws_peer     | SKIP_AWS_PEER     | aviatrix_account+AWS_VPC_ID, AWS_VPC_ID2, AWS_REGION, AWS_REGION2 |
| aviatrix_dc_extn      | SKIP_DCX          | aviatrix_account+AWS_REGION, DCX_SUBNET                      |
| aviatrix_firewall     | SKIP_FIREWALL     | aviatrix_gateway                                             |
| aviatrix_firewall_tag | SKIP_FIREWALL_TAG |                                                              |
| aviatrix_fqdn         | SKIP_FQDN         | aviatrix_gateway                                             |
| aviatrix_gateway      | SKIP_GATEWAY      | aviatrix_account+AWS_VPC_ID, AWS_REGION, AWS_VPC_NET         |
| aviatrix_site2cloud   | SKIP_S2C          | aviatrix_gateway                                             |
| aviatrix_spoke_vpc    | SKIP_SPOKE        | aviatrix_gateway                                             |
| aviatrix_transit_vpc  | SKIP_TRANSIT      | aviatrix_gateway                                             |
| aviatrix_tunnel       | SKIP_TUNNEL       | aviatrix_gateway+AWS_VPC_ID2, AWS_REGION2, AWS_VPC_NET2      |
| aviatrix_vpn_user     | SKIP_VPN_USER     | aviatrix_gateway                                             |


