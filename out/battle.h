#pragma once
#include "string"
#include "cstdint"
#include "cgHelpers/nullable.h"

using cgHelpers::Nullable;

// battle: Battle data for the game server
class battle 
{
public:
	/* Variables */
	// byArea_8: TODO: Description
	uint8_t byArea_8;

	// byKarusArea: TODO: Description
	uint8_t byKarusArea;

	// byElmoAdvantage: TODO: Description
	uint8_t byElmoAdvantage;

	// byArea_4: TODO: Description
	uint8_t byArea_4;

	// byArea_5: TODO: Description
	uint8_t byArea_5;

	// byArea_6: TODO: Description
	uint8_t byArea_6;

	// byArea_7: TODO: Description
	uint8_t byArea_7;

	// byKarusAdvantage: TODO: Description
	uint8_t byKarusAdvantage;

	// byElmoArea: TODO: Description
	uint8_t byElmoArea;

	// byNation: User Nation Identifier
	uint8_t byNation;

	// byArea_1: TODO: Description
	uint8_t byArea_1;

	// byArea_10: TODO: Description
	uint8_t byArea_10;

	// byArea_2: TODO: Description
	uint8_t byArea_2;

	// byArea_3: TODO: Description
	uint8_t byArea_3;

	// byArea_11: TODO: Description
	uint8_t byArea_11;

	// sIndex: Server Index
	int16_t sIndex;

	// byArea_9: TODO: Description
	uint8_t byArea_9;

	// strUserName: User Name
	Nullable<std::string> strUserName;

	/* Helper Functions */
	// GetTableName: Returns the database table name
	static std::string GetTableName()
	{
		return "BATTLE";
	}

	// CN_byArea_8: Returns the database column name
	static std::string CN_byArea_8()
	{
		return "byArea_8";
	}

	// CN_byKarusArea: Returns the database column name
	static std::string CN_byKarusArea()
	{
		return "byKarusArea";
	}

	// CN_byElmoAdvantage: Returns the database column name
	static std::string CN_byElmoAdvantage()
	{
		return "byElmoAdvantage";
	}

	// CN_byArea_4: Returns the database column name
	static std::string CN_byArea_4()
	{
		return "byArea_4";
	}

	// CN_byArea_5: Returns the database column name
	static std::string CN_byArea_5()
	{
		return "byArea_5";
	}

	// CN_byArea_6: Returns the database column name
	static std::string CN_byArea_6()
	{
		return "byArea_6";
	}

	// CN_byArea_7: Returns the database column name
	static std::string CN_byArea_7()
	{
		return "byArea_7";
	}

	// CN_byKarusAdvantage: Returns the database column name
	static std::string CN_byKarusAdvantage()
	{
		return "byKarusAdvantage";
	}

	// CN_byElmoArea: Returns the database column name
	static std::string CN_byElmoArea()
	{
		return "byElmoArea";
	}

	// CN_byNation: Returns the database column name
	static std::string CN_byNation()
	{
		return "byNation";
	}

	// CN_byArea_1: Returns the database column name
	static std::string CN_byArea_1()
	{
		return "byArea_1";
	}

	// CN_byArea_10: Returns the database column name
	static std::string CN_byArea_10()
	{
		return "byArea_10";
	}

	// CN_byArea_2: Returns the database column name
	static std::string CN_byArea_2()
	{
		return "byArea_2";
	}

	// CN_byArea_3: Returns the database column name
	static std::string CN_byArea_3()
	{
		return "byArea_3";
	}

	// CN_byArea_11: Returns the database column name
	static std::string CN_byArea_11()
	{
		return "byArea_11";
	}

	// CN_sIndex: Returns the database column name
	static std::string CN_sIndex()
	{
		return "sIndex";
	}

	// CN_byArea_9: Returns the database column name
	static std::string CN_byArea_9()
	{
		return "byArea_9";
	}

	// CN_strUserName: Returns the database column name
	static std::string CN_strUserName()
	{
		return "strUserName";
	}
};