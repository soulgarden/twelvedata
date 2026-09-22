package response

import "github.com/guregu/null/v6"

// ConsolidatedIncomeStatements contains income statement data from the consolidated endpoint.
type ConsolidatedIncomeStatements struct {
	IncomeStatement []ConsolidatedIncomeStatement `json:"income_statement"`
	Status          string                        `json:"status"`
}

// ConsolidatedIncomeStatement contains income statement data from the consolidated endpoint.
type ConsolidatedIncomeStatement struct {
	FiscalDate                  string                                                 `json:"fiscal_date"`
	Year                        null.Int                                               `json:"year"`
	Revenue                     ConsolidatedIncomeStatementRevenue                     `json:"revenue"`
	GrossProfit                 ConsolidatedIncomeStatementGrossProfit                 `json:"gross_profit"`
	OperatingIncome             ConsolidatedIncomeStatementOperatingIncome             `json:"operating_income"`
	NetIncome                   ConsolidatedIncomeStatementNetIncome                   `json:"net_income"`
	EarningsPerShare            ConsolidatedIncomeStatementEarningsPerShare            `json:"earnings_per_share"`
	Expenses                    ConsolidatedIncomeStatementExpenses                    `json:"expenses"`
	InterestIncomeAndExpense    ConsolidatedIncomeStatementInterestIncomeAndExpense    `json:"interest_income_and_expense"`
	OtherIncomeAndExpenses      ConsolidatedIncomeStatementOtherIncomeAndExpenses      `json:"other_income_and_expenses"`
	Taxes                       ConsolidatedIncomeStatementTaxes                       `json:"taxes"`
	DepreciationAndAmortization ConsolidatedIncomeStatementDepreciationAndAmortization `json:"depreciation_and_amortization"`
	EBITDA                      ConsolidatedIncomeStatementEBITDA                      `json:"ebitda"`
	DividendsAndShares          ConsolidatedIncomeStatementDividendsAndShares          `json:"dividends_and_shares"`
	UnusualItems                ConsolidatedIncomeStatementUnusualItems                `json:"unusual_items"`
	Depreciation                ConsolidatedIncomeStatementDepreciation                `json:"depreciation"`
	PretaxIncome                ConsolidatedIncomeStatementPretaxIncome                `json:"pretax_income"`
	SpecialIncomeCharges        ConsolidatedIncomeStatementSpecialIncomeCharges        `json:"special_income_charges"`
}

// ConsolidatedIncomeStatementRevenue contains income statement data from the consolidated endpoint.
type ConsolidatedIncomeStatementRevenue struct {
	TotalRevenue     null.Float `json:"total_revenue"`
	OperatingRevenue null.Float `json:"operating_revenue"`
}

// ConsolidatedIncomeStatementGrossProfit contains income statement data from the consolidated endpoint.
type ConsolidatedIncomeStatementGrossProfit struct {
	GrossProfitValue null.Float                               `json:"gross_profit_value"`
	CostOfRevenue    ConsolidatedIncomeStatementCostOfRevenue `json:"cost_of_revenue"`
}

// ConsolidatedIncomeStatementCostOfRevenue contains income statement data from the consolidated endpoint.
type ConsolidatedIncomeStatementCostOfRevenue struct {
	CostOfRevenueValue      null.Float `json:"cost_of_revenue_value"`
	ExciseTaxes             null.Float `json:"excise_taxes"`
	ReconciledCostOfRevenue null.Float `json:"reconciled_cost_of_revenue"`
}

// ConsolidatedIncomeStatementOperatingIncome contains income statement data from the consolidated endpoint.
type ConsolidatedIncomeStatementOperatingIncome struct {
	OperatingIncomeValue           null.Float `json:"operating_income_value"`
	TotalOperatingIncomeAsReported null.Float `json:"total_operating_income_as_reported"`
	OperatingExpense               null.Float `json:"operating_expense"`
	OtherOperatingExpenses         null.Float `json:"other_operating_expenses"`
	TotalExpenses                  null.Float `json:"total_expenses"`
}

// ConsolidatedIncomeStatementNetIncome contains income statement data from the consolidated endpoint.
type ConsolidatedIncomeStatementNetIncome struct {
	NetIncomeValue                                      null.Float `json:"net_income_value"`
	NetIncomeCommonStockholders                         null.Float `json:"net_income_common_stockholders"`
	NetIncomeIncludingNoncontrollingInterests           null.Float `json:"net_income_including_noncontrolling_interests"`
	NetIncomeFromTaxLossCarryforward                    null.Float `json:"net_income_from_tax_loss_carryforward"`
	NetIncomeExtraordinary                              null.Float `json:"net_income_extraordinary"`
	NetIncomeDiscontinuousOperations                    null.Float `json:"net_income_discontinuous_operations"`
	NetIncomeContinuousOperations                       null.Float `json:"net_income_continuous_operations"`
	NetIncomeFromContinuingOperationNetMinorityInterest null.Float `json:"net_income_from_continuing_operation_net_minority_interest"`
	NetIncomeFromContinuingAndDiscontinuedOperation     null.Float `json:"net_income_from_continuing_and_discontinued_operation"`
	NormalizedIncome                                    null.Float `json:"normalized_income"`
	MinorityInterests                                   null.Float `json:"minority_interests"`
}

// ConsolidatedIncomeStatementEarningsPerShare contains income statement data from the consolidated endpoint.
type ConsolidatedIncomeStatementEarningsPerShare struct {
	DilutedEPS                          null.Float `json:"diluted_eps"`
	BasicEPS                            null.Float `json:"basic_eps"`
	ContinuingAndDiscontinuedDilutedEPS null.Float `json:"continuing_and_discontinued_diluted_eps"`
	ContinuingAndDiscontinuedBasicEPS   null.Float `json:"continuing_and_discontinued_basic_eps"`
	NormalizedDilutedEPS                null.Float `json:"normalized_diluted_eps"`
	NormalizedBasicEPS                  null.Float `json:"normalized_basic_eps"`
	ReportedNormalizedDilutedEPS        null.Float `json:"reported_normalized_diluted_eps"`
	ReportedNormalizedBasicEPS          null.Float `json:"reported_normalized_basic_eps"`
	DilutedEPSOtherGainsLosses          null.Float `json:"diluted_eps_other_gains_losses"`
	TaxLossCarryforwardDilutedEPS       null.Float `json:"tax_loss_carryforward_diluted_eps"`
	DilutedAccountingChange             null.Float `json:"diluted_accounting_change"`
	DilutedExtraordinary                null.Float `json:"diluted_extraordinary"`
	DilutedDiscontinuousOperations      null.Float `json:"diluted_discontinuous_operations"`
	DilutedContinuousOperations         null.Float `json:"diluted_continuous_operations"`
	BasicEPSOtherGainsLosses            null.Float `json:"basic_eps_other_gains_losses"`
	TaxLossCarryforwardBasicEPS         null.Float `json:"tax_loss_carryforward_basic_eps"`
	BasicAccountingChange               null.Float `json:"basic_accounting_change"`
	BasicExtraordinary                  null.Float `json:"basic_extraordinary"`
	BasicDiscontinuousOperations        null.Float `json:"basic_discontinuous_operations"`
	BasicContinuousOperations           null.Float `json:"basic_continuous_operations"`
	DilutedNiAvailToCommonStockholders  null.Float `json:"diluted_ni_avail_to_common_stockholders"`
	AverageDilutionEarnings             null.Float `json:"average_dilution_earnings"`
}

// ConsolidatedIncomeStatementExpenses contains income statement data from the consolidated endpoint.
type ConsolidatedIncomeStatementExpenses struct {
	TotalExpenses                                    null.Float `json:"total_expenses"`
	SellingGeneralAndAdministrationExpense           null.Float `json:"selling_general_and_administration_expense"`
	SellingAndMarketingExpense                       null.Float `json:"selling_and_marketing_expense"`
	GeneralAndAdministrativeExpense                  null.Float `json:"general_and_administrative_expense"`
	OtherGeneralAndAdministrativeExpense             null.Float `json:"other_general_and_administrative_expense"`
	DepreciationAmortizationDepletionIncomeStatement null.Float `json:"depreciation_amortization_depletion_income_statement"`
	ResearchAndDevelopmentExpense                    null.Float `json:"research_and_development_expense"`
	InsuranceAndClaimsExpense                        null.Float `json:"insurance_and_claims_expense"`
	RentAndLandingFees                               null.Float `json:"rent_and_landing_fees"`
	SalariesAndWagesExpense                          null.Float `json:"salaries_and_wages_expense"`
	RentExpenseSupplemental                          null.Float `json:"rent_expense_supplemental"`
	ProvisionForDoubtfulAccounts                     null.Float `json:"provision_for_doubtful_accounts"`
}

// ConsolidatedIncomeStatementInterestIncomeAndExpense contains income statement data from the consolidated endpoint.
type ConsolidatedIncomeStatementInterestIncomeAndExpense struct {
	InterestIncome                       null.Float `json:"interest_income"`
	InterestExpense                      null.Float `json:"interest_expense"`
	NetInterestIncome                    null.Float `json:"net_interest_income"`
	NetNonOperatingInterestIncomeExpense null.Float `json:"net_non_operating_interest_income_expense"`
	InterestExpenseNonOperating          null.Float `json:"interest_expense_non_operating"`
	InterestIncomeNonOperating           null.Float `json:"interest_income_non_operating"`
}

// ConsolidatedIncomeStatementOtherIncomeAndExpenses contains income statement data from the consolidated endpoint.
type ConsolidatedIncomeStatementOtherIncomeAndExpenses struct {
	OtherIncomeExpense                 null.Float `json:"other_income_expense"`
	OtherNonOperatingIncomeExpenses    null.Float `json:"other_non_operating_income_expenses"`
	SpecialIncomeCharges               null.Float `json:"special_income_charges"`
	GainOnSaleOfPPE                    null.Float `json:"gain_on_sale_of_ppe"`
	GainOnSaleOfBusiness               null.Float `json:"gain_on_sale_of_business"`
	GainOnSaleOfSecurity               null.Float `json:"gain_on_sale_of_security"`
	OtherSpecialCharges                null.Float `json:"other_special_charges"`
	WriteOff                           null.Float `json:"write_off"`
	ImpairmentOfCapitalAssets          null.Float `json:"impairment_of_capital_assets"`
	RestructuringAndMergerAcquisition  null.Float `json:"restructuring_and_merger_acquisition"`
	SecuritiesAmortization             null.Float `json:"securities_amortization"`
	EarningsFromEquityInterest         null.Float `json:"earnings_from_equity_interest"`
	EarningsFromEquityInterestNetOfTax null.Float `json:"earnings_from_equity_interest_net_of_tax"`
	TotalOtherFinanceCost              null.Float `json:"total_other_finance_cost"`
}

// ConsolidatedIncomeStatementTaxes contains income statement data from the consolidated endpoint.
type ConsolidatedIncomeStatementTaxes struct {
	TaxProvision            null.Float `json:"tax_provision"`
	TaxEffectOfUnusualItems null.Float `json:"tax_effect_of_unusual_items"`
	TaxRateForCalculations  null.Float `json:"tax_rate_for_calculations"`
	OtherTaxes              null.Float `json:"other_taxes"`
}

// ConsolidatedIncomeStatementDepreciationAndAmortization contains income statement data from the consolidated endpoint.
type ConsolidatedIncomeStatementDepreciationAndAmortization struct {
	DepreciationAmortizationDepletion            null.Float `json:"depreciation_amortization_depletion"`
	AmortizationOfIntangibles                    null.Float `json:"amortization_of_intangibles"`
	Depreciation                                 null.Float `json:"depreciation"`
	Amortization                                 null.Float `json:"amortization"`
	Depletion                                    null.Float `json:"depletion"`
	DepreciationAndAmortizationInIncomeStatement null.Float `json:"depreciation_and_amortization_in_income_statement"`
}

// ConsolidatedIncomeStatementEBITDA contains income statement data from the consolidated endpoint.
type ConsolidatedIncomeStatementEBITDA struct {
	EBITDAValue           null.Float `json:"ebitda_value"`
	NormalizedEBITDAValue null.Float `json:"normalized_ebitda_value"`
	EBITValue             null.Float `json:"ebit_value"`
}

// ConsolidatedIncomeStatementDividendsAndShares contains income statement data from the consolidated endpoint.
type ConsolidatedIncomeStatementDividendsAndShares struct {
	DividendPerShare                 null.Float `json:"dividend_per_share"`
	DilutedAverageShares             null.Float `json:"diluted_average_shares"`
	BasicAverageShares               null.Float `json:"basic_average_shares"`
	PreferredStockDividends          null.Float `json:"preferred_stock_dividends"`
	OtherUnderPreferredStockDividend null.Float `json:"other_under_preferred_stock_dividend"`
}

// ConsolidatedIncomeStatementUnusualItems contains income statement data from the consolidated endpoint.
type ConsolidatedIncomeStatementUnusualItems struct {
	TotalUnusualItems                  null.Float `json:"total_unusual_items"`
	TotalUnusualItemsExcludingGoodwill null.Float `json:"total_unusual_items_excluding_goodwill"`
}

// ConsolidatedIncomeStatementDepreciation contains income statement data from the consolidated endpoint.
type ConsolidatedIncomeStatementDepreciation struct {
	ReconciledDepreciation null.Float `json:"reconciled_depreciation"`
}

// ConsolidatedIncomeStatementPretaxIncome contains income statement data from the consolidated endpoint.
type ConsolidatedIncomeStatementPretaxIncome struct {
	PretaxIncomeValue null.Float `json:"pretax_income_value"`
}

// ConsolidatedIncomeStatementSpecialIncomeCharges contains income statement data from the consolidated endpoint.
type ConsolidatedIncomeStatementSpecialIncomeCharges struct {
	SpecialIncomeChargesValue null.Float `json:"special_income_charges_value"`
}
