#pragma once
#include "cstdint"
#include "cgHelpers/nullable.h"
#include "string"

using cgHelpers::Nullable;

// account_char: Represents the relationship between accounts and characters
class account_char 
{
public:
	/* Variables */
	// bNation: Nation Identifier
	uint8_t bNation;

	// bCharNum: Number of characters associated with the account
	uint8_t bCharNum;

	// strCharID1: First Character Name
	Nullable<std::string> strCharID1;

	// strCharID2: Second Character Name
	Nullable<std::string> strCharID2;

	// strCharID3: Third Character Name
	Nullable<std::string> strCharID3;

	// strAccountID: Account Identifier
	std::string strAccountID;

	/* Helper Functions */
	// GetTableName: Returns the database table name
	static std::string GetTableName()
	{
		return "ACCOUNT_CHAR";
	}

	// CN_bNation: Returns the database column name
	static std::string CN_bNation()
	{
		return "bNation";
	}

	// CN_bCharNum: Returns the database column name
	static std::string CN_bCharNum()
	{
		return "bCharNum";
	}

	// CN_strCharID1: Returns the database column name
	static std::string CN_strCharID1()
	{
		return "strCharID1";
	}

	// CN_strCharID2: Returns the database column name
	static std::string CN_strCharID2()
	{
		return "strCharID2";
	}

	// CN_strCharID3: Returns the database column name
	static std::string CN_strCharID3()
	{
		return "strCharID3";
	}

	// CN_strAccountID: Returns the database column name
	static std::string CN_strAccountID()
	{
		return "strAccountID";
	}
};