<?php

if(isset($_GET["page"]))
{
	if ($_GET["page"] === "statement") {
		if (isset($_GET["do"]) && $_GET["do"] === "search")
			include "customer_account_statement_date.php";
		else
			include "customer_account_statement.php";
	}
	
	elseif ($_GET["page"] === "mini-statement")
		include "customer_mini_statement.php";
	
	elseif ($_GET["page"] === "issue") {
		if (isset($_GET["do"]) && $_GET["do"] === "process")
			include "customer_issue_atm_process.php";
		else
			include "customer_issue_atm.php";
	}
	
	elseif ($_GET["page"] === "personal")
		include "customer_personal_details.php";

	elseif ($_GET["page"] === "changepass")
		include "change_password_customer.php";

	elseif ($_GET["page"] === "beneficiary") {
		if (isset($_GET["do"])) {
			if ($_GET["do"] === "add")
				include "add_beneficiary.php";
			elseif ($_GET["do"] === "process")
				include "add_beneficiary_process.php";
			elseif ($_GET["do"] === "delete")
				include "delete_beneficiary.php";
			else
				include "display_beneficiary.php";
		} else
			include "display_beneficiary.php";

	}

	elseif ($_GET["page"] === "transfer") {
		if (isset($_GET["do"]) && $_GET["do"] === "process")
			include "customer_transfer_process.php";
		else
			include "customer_transfer.php";
	}

	else 
		include "customer_account_summary.php";

} else {
	include "customer_account_summary.php";
}


?>