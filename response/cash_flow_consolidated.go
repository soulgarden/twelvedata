package response

import "github.com/guregu/null/v6"

// ConsolidatedCashFlows contains cash flow data from the consolidated endpoint.
type ConsolidatedCashFlows struct {
	CashFlow []ConsolidatedCashFlow `json:"cash_flow"`
	Status   string                 `json:"status"`
}

// ConsolidatedCashFlow contains cash flow data from the consolidated endpoint.
type ConsolidatedCashFlow struct {
	FiscalDate                      string                                      `json:"fiscal_date"`
	Year                            null.Int                                    `json:"year"`
	CashFlowFromOperatingActivities ConsolidatedCashFlowOperatingActivities     `json:"cash_flow_from_operating_activities"`
	CashFlowFromInvestingActivities ConsolidatedCashFlowInvestingActivities     `json:"cash_flow_from_investing_activities"`
	CashFlowFromFinancingActivities ConsolidatedCashFlowFinancingActivities     `json:"cash_flow_from_financing_activities"`
	SupplementalData                ConsolidatedCashFlowSupplementalData        `json:"supplemental_data"`
	ForeignAndDomesticSales         ConsolidatedCashFlowForeignAndDomesticSales `json:"foreign_and_domestic_sales"`
	CashPosition                    ConsolidatedCashPosition                    `json:"cash_position"`
	DirectMethodCashFlow            ConsolidatedCashFlowDirectMethodCashFlow    `json:"direct_method_cash_flow"`
}

// ConsolidatedCashFlowOperatingActivities contains cash flow data from the consolidated endpoint.
type ConsolidatedCashFlowOperatingActivities struct {
	NetIncomeFromContinuingOperations            null.Float `json:"net_income_from_continuing_operations"`
	OperatingCashFlow                            null.Float `json:"operating_cash_flow"`
	CashFlowFromContinuingOperatingActivities    null.Float `json:"cash_flow_from_continuing_operating_activities"`
	CashFromDiscontinuedOperatingActivities      null.Float `json:"cash_from_discontinued_operating_activities"`
	CashFlowFromDiscontinuedOperation            null.Float `json:"cash_flow_from_discontinued_operation"`
	FreeCashFlow                                 null.Float `json:"free_cash_flow"`
	CashFlowsFromUsedInOperatingActivitiesDirect null.Float `json:"cash_flows_from_used_in_operating_activities_direct"`
	TaxesRefundPaid                              null.Float `json:"taxes_refund_paid"`
	TaxesRefundPaidDirect                        null.Float `json:"taxes_refund_paid_direct"`
	InterestReceived                             null.Float `json:"interest_received"`
	InterestReceivedDirect                       null.Float `json:"interest_received_direct"`
	InterestPaid                                 null.Float `json:"interest_paid"`
	InterestPaidDirect                           null.Float `json:"interest_paid_direct"`
	DividendReceived                             null.Float `json:"dividend_received"`
	DividendReceivedDirect                       null.Float `json:"dividend_received_direct"`
	DividendPaid                                 null.Float `json:"dividend_paid"`
	DividendPaidDirect                           null.Float `json:"dividend_paid_direct"`
	ChangeInWorkingCapital                       null.Float `json:"change_in_working_capital"`
	ChangeInOtherWorkingCapital                  null.Float `json:"change_in_other_working_capital"`
	ChangeInReceivables                          null.Float `json:"change_in_receivables"`
	ChangesInAccountReceivables                  null.Float `json:"changes_in_account_receivables"`
	ChangeInPayablesAndAccruedExpense            null.Float `json:"change_in_payables_and_accrued_expense"`
	ChangeInAccruedExpense                       null.Float `json:"change_in_accrued_expense"`
	ChangeInPayable                              null.Float `json:"change_in_payable"`
	ChangeInDividendPayable                      null.Float `json:"change_in_dividend_payable"`
	ChangeInAccountPayable                       null.Float `json:"change_in_account_payable"`
	ChangeInTaxPayable                           null.Float `json:"change_in_tax_payable"`
	ChangeInIncomeTaxPayable                     null.Float `json:"change_in_income_tax_payable"`
	ChangeInInterestPayable                      null.Float `json:"change_in_interest_payable"`
	ChangeInOtherCurrentLiabilities              null.Float `json:"change_in_other_current_liabilities"`
	ChangeInOtherCurrentAssets                   null.Float `json:"change_in_other_current_assets"`
	ChangeInInventory                            null.Float `json:"change_in_inventory"`
	ChangeInPrepaidAssets                        null.Float `json:"change_in_prepaid_assets"`
	OtherNonCashItems                            null.Float `json:"other_non_cash_items"`
	ExcessTaxBenefitFromStockBasedCompensation   null.Float `json:"excess_tax_benefit_from_stock_based_compensation"`
	StockBasedCompensation                       null.Float `json:"stock_based_compensation"`
	UnrealizedGainLossOnInvestmentSecurities     null.Float `json:"unrealized_gain_loss_on_investment_securities"`
	ProvisionAndWriteOffOfAssets                 null.Float `json:"provision_and_write_off_of_assets"`
	AssetImpairmentCharge                        null.Float `json:"asset_impairment_charge"`
	AmortizationOfSecurities                     null.Float `json:"amortization_of_securities"`
	DeferredTax                                  null.Float `json:"deferred_tax"`
	DeferredIncomeTax                            null.Float `json:"deferred_income_tax"`
	DepreciationAmortizationDepletion            null.Float `json:"depreciation_amortization_depletion"`
	Depletion                                    null.Float `json:"depletion"`
	DepreciationAndAmortization                  null.Float `json:"depreciation_and_amortization"`
	AmortizationCashFlow                         null.Float `json:"amortization_cash_flow"`
	AmortizationOfIntangibles                    null.Float `json:"amortization_of_intangibles"`
	Depreciation                                 null.Float `json:"depreciation"`
	OperatingGainsLosses                         null.Float `json:"operating_gains_losses"`
	PensionAndEmployeeBenefitExpense             null.Float `json:"pension_and_employee_benefit_expense"`
	EarningsLossesFromEquityInvestments          null.Float `json:"earnings_losses_from_equity_investments"`
	GainLossOnInvestmentSecurities               null.Float `json:"gain_loss_on_investment_securities"`
	NetForeignCurrencyExchangeGainLoss           null.Float `json:"net_foreign_currency_exchange_gain_loss"`
	GainLossOnSaleOfPPE                          null.Float `json:"gain_loss_on_sale_of_ppe"`
	GainLossOnSaleOfBusiness                     null.Float `json:"gain_loss_on_sale_of_business"`
}

// ConsolidatedCashFlowInvestingActivities contains cash flow data from the consolidated endpoint.
type ConsolidatedCashFlowInvestingActivities struct {
	InvestingCashFlow                         null.Float `json:"investing_cash_flow"`
	CashFlowFromContinuingInvestingActivities null.Float `json:"cash_flow_from_continuing_investing_activities"`
	CashFromDiscontinuedInvestingActivities   null.Float `json:"cash_from_discontinued_investing_activities"`
	NetOtherInvestingChanges                  null.Float `json:"net_other_investing_changes"`
	InterestReceivedCfi                       null.Float `json:"interest_received_cfi"`
	DividendsReceivedCfi                      null.Float `json:"dividends_received_cfi"`
	NetInvestmentPurchaseAndSale              null.Float `json:"net_investment_purchase_and_sale"`
	SaleOfInvestment                          null.Float `json:"sale_of_investment"`
	PurchaseOfInvestment                      null.Float `json:"purchase_of_investment"`
	NetInvestmentPropertiesPurchaseAndSale    null.Float `json:"net_investment_properties_purchase_and_sale"`
	SaleOfInvestmentProperties                null.Float `json:"sale_of_investment_properties"`
	PurchaseOfInvestmentProperties            null.Float `json:"purchase_of_investment_properties"`
	NetBusinessPurchaseAndSale                null.Float `json:"net_business_purchase_and_sale"`
	SaleOfBusiness                            null.Float `json:"sale_of_business"`
	PurchaseOfBusiness                        null.Float `json:"purchase_of_business"`
	NetIntangiblesPurchaseAndSale             null.Float `json:"net_intangibles_purchase_and_sale"`
	SaleOfIntangibles                         null.Float `json:"sale_of_intangibles"`
	PurchaseOfIntangibles                     null.Float `json:"purchase_of_intangibles"`
	NetPPEPurchaseAndSale                     null.Float `json:"net_ppe_purchase_and_sale"`
	SaleOfPPE                                 null.Float `json:"sale_of_ppe"`
	PurchaseOfPPE                             null.Float `json:"purchase_of_ppe"`
	CapitalExpenditureReported                null.Float `json:"capital_expenditure_reported"`
	CapitalExpenditure                        null.Float `json:"capital_expenditure"`
}

// ConsolidatedCashFlowFinancingActivities contains cash flow data from the consolidated endpoint.
type ConsolidatedCashFlowFinancingActivities struct {
	FinancingCashFlow                         null.Float `json:"financing_cash_flow"`
	CashFlowFromContinuingFinancingActivities null.Float `json:"cash_flow_from_continuing_financing_activities"`
	CashFromDiscontinuedFinancingActivities   null.Float `json:"cash_from_discontinued_financing_activities"`
	NetOtherFinancingCharges                  null.Float `json:"net_other_financing_charges"`
	InterestPaidCff                           null.Float `json:"interest_paid_cff"`
	ProceedsFromStockOptionExercised          null.Float `json:"proceeds_from_stock_option_exercised"`
	CashDividendsPaid                         null.Float `json:"cash_dividends_paid"`
	PreferredStockDividendPaid                null.Float `json:"preferred_stock_dividend_paid"`
	CommonStockDividendPaid                   null.Float `json:"common_stock_dividend_paid"`
	NetPreferredStockIssuance                 null.Float `json:"net_preferred_stock_issuance"`
	PreferredStockPayments                    null.Float `json:"preferred_stock_payments"`
	PreferredStockIssuance                    null.Float `json:"preferred_stock_issuance"`
	NetCommonStockIssuance                    null.Float `json:"net_common_stock_issuance"`
	CommonStockPayments                       null.Float `json:"common_stock_payments"`
	CommonStockIssuance                       null.Float `json:"common_stock_issuance"`
	RepurchaseOfCapitalStock                  null.Float `json:"repurchase_of_capital_stock"`
	NetIssuancePaymentsOfDebt                 null.Float `json:"net_issuance_payments_of_debt"`
	NetShortTermDebtIssuance                  null.Float `json:"net_short_term_debt_issuance"`
	ShortTermDebtPayments                     null.Float `json:"short_term_debt_payments"`
	ShortTermDebtIssuance                     null.Float `json:"short_term_debt_issuance"`
	NetLongTermDebtIssuance                   null.Float `json:"net_long_term_debt_issuance"`
	LongTermDebtPayments                      null.Float `json:"long_term_debt_payments"`
	LongTermDebtIssuance                      null.Float `json:"long_term_debt_issuance"`
	IssuanceOfDebt                            null.Float `json:"issuance_of_debt"`
	RepaymentOfDebt                           null.Float `json:"repayment_of_debt"`
	IssuanceOfCapitalStock                    null.Float `json:"issuance_of_capital_stock"`
}

// ConsolidatedCashFlowSupplementalData contains cash flow data from the consolidated endpoint.
type ConsolidatedCashFlowSupplementalData struct {
	InterestPaidSupplementalData  null.Float `json:"interest_paid_supplemental_data"`
	IncomeTaxPaidSupplementalData null.Float `json:"income_tax_paid_supplemental_data"`
}

// ConsolidatedCashFlowForeignAndDomesticSales contains cash flow data from the consolidated endpoint.
type ConsolidatedCashFlowForeignAndDomesticSales struct {
	ForeignSales                 null.Float `json:"foreign_sales"`
	DomesticSales                null.Float `json:"domestic_sales"`
	AdjustedGeographySegmentData null.Float `json:"adjusted_geography_segment_data"`
}

// ConsolidatedCashPosition contains cash flow data from the consolidated endpoint.
type ConsolidatedCashPosition struct {
	BeginningCashPosition                  null.Float `json:"beginning_cash_position"`
	EndCashPosition                        null.Float `json:"end_cash_position"`
	ChangesInCash                          null.Float `json:"changes_in_cash"`
	OtherCashAdjustmentOutsideChangeInCash null.Float `json:"other_cash_adjustment_outside_change_in_cash"`
	OtherCashAdjustmentInsideChangeInCash  null.Float `json:"other_cash_adjustment_inside_change_in_cash"`
	EffectOfExchangeRateChanges            null.Float `json:"effect_of_exchange_rate_changes"`
}

// ConsolidatedCashFlowDirectMethodCashFlow contains cash flow data from the consolidated endpoint.
type ConsolidatedCashFlowDirectMethodCashFlow struct {
	ClassesOfCashReceiptsFromOperatingActivities null.Float `json:"classes_of_cash_receipts_from_operating_activities"`
	OtherCashReceiptsFromOperatingActivities     null.Float `json:"other_cash_receipts_from_operating_activities"`
	ReceiptsFromGovernmentGrants                 null.Float `json:"receipts_from_government_grants"`
	ReceiptsFromCustomers                        null.Float `json:"receipts_from_customers"`
	ClassesOfCashPayments                        null.Float `json:"classes_of_cash_payments"`
	OtherCashPaymentsFromOperatingActivities     null.Float `json:"other_cash_payments_from_operating_activities"`
	PaymentsOnBehalfOfEmployees                  null.Float `json:"payments_on_behalf_of_employees"`
	PaymentsToSuppliersForGoodsAndServices       null.Float `json:"payments_to_suppliers_for_goods_and_services"`
}
