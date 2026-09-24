package response

import "github.com/guregu/null/v6"

// ConsolidatedBalanceSheets contains balance sheet data from the consolidated endpoint.
type ConsolidatedBalanceSheets struct {
	BalanceSheet []ConsolidatedBalanceSheet `json:"balance_sheet"`
	Status       string                     `json:"status"`
}

// ConsolidatedBalanceSheet contains balance sheet data from the consolidated endpoint.
type ConsolidatedBalanceSheet struct {
	FiscalDate string                         `json:"fiscal_date"`
	Assets     ConsolidatedBalanceSheetAssets `json:"assets"`
}

// ConsolidatedBalanceSheetAssets contains balance sheet data from the consolidated endpoint.
type ConsolidatedBalanceSheetAssets struct {
	TotalAssets      null.Float                               `json:"total_assets"`
	CurrentAssets    ConsolidatedBalanceSheetCurrentAssets    `json:"current_assets"`
	NonCurrentAssets ConsolidatedBalanceSheetNonCurrentAssets `json:"non_current_assets"`
	Liabilities      ConsolidatedBalanceSheetLiabilities      `json:"liabilities"`
}

// ConsolidatedBalanceSheetCurrentAssets contains balance sheet data from the consolidated endpoint.
type ConsolidatedBalanceSheetCurrentAssets struct {
	TotalCurrentAssets                         null.Float                          `json:"total_current_assets"`
	CashCashEquivalentsAndShortTermInvestments null.Float                          `json:"cash_cash_equivalents_and_short_term_investments"`
	CashAndCashEquivalents                     null.Float                          `json:"cash_and_cash_equivalents"`
	CashEquivalents                            null.Float                          `json:"cash_equivalents"`
	CashFinancial                              null.Float                          `json:"cash_financial"`
	OtherShortTermInvestments                  null.Float                          `json:"other_short_term_investments"`
	RestrictedCash                             null.Float                          `json:"restricted_cash"`
	Receivables                                ConsolidatedBalanceSheetReceivables `json:"receivables"`
	Inventory                                  ConsolidatedBalanceSheetInventory   `json:"inventory"`
	PrepaidAssets                              null.Float                          `json:"prepaid_assets"`
	CurrentDeferredAssets                      null.Float                          `json:"current_deferred_assets"`
	CurrentDeferredTaxesAssets                 null.Float                          `json:"current_deferred_taxes_assets"`
	AssetsHeldForSaleCurrent                   null.Float                          `json:"assets_held_for_sale_current"`
	HedgingAssetsCurrent                       null.Float                          `json:"hedging_assets_current"`
	OtherCurrentAssets                         null.Float                          `json:"other_current_assets"`
}

// ConsolidatedBalanceSheetReceivables contains balance sheet data from the consolidated endpoint.
type ConsolidatedBalanceSheetReceivables struct {
	TotalReceivables                       null.Float `json:"total_receivables"`
	AccountsReceivable                     null.Float `json:"accounts_receivable"`
	GrossAccountsReceivable                null.Float `json:"gross_accounts_receivable"`
	AllowanceForDoubtfulAccountsReceivable null.Float `json:"allowance_for_doubtful_accounts_receivable"`
	ReceivablesAdjustmentsAllowances       null.Float `json:"receivables_adjustments_allowances"`
	OtherReceivables                       null.Float `json:"other_receivables"`
	DueFromRelatedPartiesCurrent           null.Float `json:"due_from_related_parties_current"`
	TaxesReceivable                        null.Float `json:"taxes_receivable"`
	AccruedInterestReceivable              null.Float `json:"accrued_interest_receivable"`
	NotesReceivable                        null.Float `json:"notes_receivable"`
	LoansReceivable                        null.Float `json:"loans_receivable"`
}

// ConsolidatedBalanceSheetInventory contains balance sheet data from the consolidated endpoint.
type ConsolidatedBalanceSheetInventory struct {
	TotalInventory                   null.Float `json:"total_inventory"`
	InventoriesAdjustmentsAllowances null.Float `json:"inventories_adjustments_allowances"`
	OtherInventories                 null.Float `json:"other_inventories"`
	FinishedGoods                    null.Float `json:"finished_goods"`
	WorkInProcess                    null.Float `json:"work_in_process"`
	RawMaterials                     null.Float `json:"raw_materials"`
}

// ConsolidatedBalanceSheetNonCurrentAssets contains balance sheet data from the consolidated endpoint.
type ConsolidatedBalanceSheetNonCurrentAssets struct {
	TotalNonCurrentAssets                                        null.Float                                               `json:"total_non_current_assets"`
	FinancialAssets                                              null.Float                                               `json:"financial_assets"`
	InvestmentsAndAdvances                                       null.Float                                               `json:"investments_and_advances"`
	OtherInvestments                                             null.Float                                               `json:"other_investments"`
	InvestmentInFinancialAssets                                  null.Float                                               `json:"investment_in_financial_assets"`
	HeldToMaturitySecurities                                     null.Float                                               `json:"held_to_maturity_securities"`
	AvailableForSaleSecurities                                   null.Float                                               `json:"available_for_sale_securities"`
	FinancialAssetsDesignatedAsFairValueThroughProfitOrLossTotal null.Float                                               `json:"financial_assets_designated_as_fair_value_through_profit_or_loss_total"`
	TradingSecurities                                            null.Float                                               `json:"trading_securities"`
	LongTermEquityInvestment                                     null.Float                                               `json:"long_term_equity_investment"`
	InvestmentsInJointVenturesAtCost                             null.Float                                               `json:"investments_in_joint_ventures_at_cost"`
	InvestmentsInOtherVenturesUnderEquityMethod                  null.Float                                               `json:"investments_in_other_ventures_under_equity_method"`
	InvestmentsInAssociatesAtCost                                null.Float                                               `json:"investments_in_associates_at_cost"`
	InvestmentsInSubsidiariesAtCost                              null.Float                                               `json:"investments_in_subsidiaries_at_cost"`
	InvestmentProperties                                         null.Float                                               `json:"investment_properties"`
	GoodwillAndOtherIntangibleAssets                             ConsolidatedBalanceSheetGoodwillAndOtherIntangibleAssets `json:"goodwill_and_other_intangible_assets"`
	NetPPE                                                       null.Float                                               `json:"net_ppe"`
	GrossPPE                                                     null.Float                                               `json:"gross_ppe"`
	AccumulatedDepreciation                                      null.Float                                               `json:"accumulated_depreciation"`
	Leases                                                       null.Float                                               `json:"leases"`
	ConstructionInProgress                                       null.Float                                               `json:"construction_in_progress"`
	OtherProperties                                              null.Float                                               `json:"other_properties"`
	MachineryFurnitureEquipment                                  null.Float                                               `json:"machinery_furniture_equipment"`
	BuildingsAndImprovements                                     null.Float                                               `json:"buildings_and_improvements"`
	LandAndImprovements                                          null.Float                                               `json:"land_and_improvements"`
	Properties                                                   null.Float                                               `json:"properties"`
	NonCurrentAccountsReceivable                                 null.Float                                               `json:"non_current_accounts_receivable"`
	NonCurrentNoteReceivables                                    null.Float                                               `json:"non_current_note_receivables"`
	DueFromRelatedPartiesNonCurrent                              null.Float                                               `json:"due_from_related_parties_non_current"`
	NonCurrentPrepaidAssets                                      null.Float                                               `json:"non_current_prepaid_assets"`
	NonCurrentDeferredAssets                                     null.Float                                               `json:"non_current_deferred_assets"`
	NonCurrentDeferredTaxesAssets                                null.Float                                               `json:"non_current_deferred_taxes_assets"`
	DefinedPensionBenefit                                        null.Float                                               `json:"defined_pension_benefit"`
	OtherNonCurrentAssets                                        null.Float                                               `json:"other_non_current_assets"`
}

// ConsolidatedBalanceSheetGoodwillAndOtherIntangibleAssets contains balance sheet data from the consolidated endpoint.
type ConsolidatedBalanceSheetGoodwillAndOtherIntangibleAssets struct {
	Goodwill                         null.Float `json:"goodwill"`
	OtherIntangibleAssets            null.Float `json:"other_intangible_assets"`
	TotalGoodwillAndIntangibleAssets null.Float `json:"total_goodwill_and_intangible_assets"`
}

// ConsolidatedBalanceSheetLiabilities contains balance sheet data from the consolidated endpoint.
type ConsolidatedBalanceSheetLiabilities struct {
	TotalLiabilitiesNetMinorityInterest null.Float                                    `json:"total_liabilities_net_minority_interest"`
	CurrentLiabilities                  ConsolidatedBalanceSheetCurrentLiabilities    `json:"current_liabilities"`
	NonCurrentLiabilities               ConsolidatedBalanceSheetNonCurrentLiabilities `json:"non_current_liabilities"`
	Equity                              ConsolidatedBalanceSheetEquity                `json:"equity"`
}

// ConsolidatedBalanceSheetCurrentLiabilities contains balance sheet data from the consolidated endpoint.
type ConsolidatedBalanceSheetCurrentLiabilities struct {
	TotalCurrentLiabilities                          null.Float                                         `json:"total_current_liabilities"`
	CurrentDebtAndCapitalLeaseObligation             null.Float                                         `json:"current_debt_and_capital_lease_obligation"`
	CurrentDebt                                      null.Float                                         `json:"current_debt"`
	CurrentCapitalLeaseObligation                    null.Float                                         `json:"current_capital_lease_obligation"`
	OtherCurrentBorrowings                           null.Float                                         `json:"other_current_borrowings"`
	LineOfCredit                                     null.Float                                         `json:"line_of_credit"`
	CommercialPaper                                  null.Float                                         `json:"commercial_paper"`
	CurrentNotesPayable                              null.Float                                         `json:"current_notes_payable"`
	CurrentProvisions                                null.Float                                         `json:"current_provisions"`
	PayablesAndAccruedExpenses                       ConsolidatedBalanceSheetPayablesAndAccruedExpenses `json:"payables_and_accrued_expenses"`
	PensionAndOtherPostRetirementBenefitPlansCurrent null.Float                                         `json:"pension_and_other_post_retirement_benefit_plans_current"`
	EmployeeBenefits                                 null.Float                                         `json:"employee_benefits"`
	CurrentDeferredLiabilities                       null.Float                                         `json:"current_deferred_liabilities"`
	CurrentDeferredRevenue                           null.Float                                         `json:"current_deferred_revenue"`
	CurrentDeferredTaxesLiabilities                  null.Float                                         `json:"current_deferred_taxes_liabilities"`
	OtherCurrentLiabilities                          null.Float                                         `json:"other_current_liabilities"`
	LiabilitiesHeldForSaleNonCurrent                 null.Float                                         `json:"liabilities_held_for_sale_non_current"`
}

// ConsolidatedBalanceSheetPayablesAndAccruedExpenses contains balance sheet data from the consolidated endpoint.
type ConsolidatedBalanceSheetPayablesAndAccruedExpenses struct {
	TotalPayablesAndAccruedExpenses null.Float `json:"total_payables_and_accrued_expenses"`
	AccountsPayable                 null.Float `json:"accounts_payable"`
	CurrentAccruedExpenses          null.Float `json:"current_accrued_expenses"`
	InterestPayable                 null.Float `json:"interest_payable"`
	Payables                        null.Float `json:"payables"`
	OtherPayable                    null.Float `json:"other_payable"`
	DueToRelatedPartiesCurrent      null.Float `json:"due_to_related_parties_current"`
	DividendsPayable                null.Float `json:"dividends_payable"`
	TotalTaxPayable                 null.Float `json:"total_tax_payable"`
	IncomeTaxPayable                null.Float `json:"income_tax_payable"`
}

// ConsolidatedBalanceSheetNonCurrentLiabilities contains balance sheet data from the consolidated endpoint.
type ConsolidatedBalanceSheetNonCurrentLiabilities struct {
	TotalNonCurrentLiabilitiesNetMinorityInterest       null.Float                                                    `json:"total_non_current_liabilities_net_minority_interest"`
	LongTermDebtAndCapitalLeaseObligation               ConsolidatedBalanceSheetLongTermDebtAndCapitalLeaseObligation `json:"long_term_debt_and_capital_lease_obligation"`
	LongTermProvisions                                  null.Float                                                    `json:"long_term_provisions"`
	NonCurrentPensionAndOtherPostretirementBenefitPlans null.Float                                                    `json:"non_current_pension_and_other_postretirement_benefit_plans"`
	NonCurrentAccruedExpenses                           null.Float                                                    `json:"non_current_accrued_expenses"`
	DueToRelatedPartiesNonCurrent                       null.Float                                                    `json:"due_to_related_parties_non_current"`
	TradeAndOtherPayablesNonCurrent                     null.Float                                                    `json:"trade_and_other_payables_non_current"`
	NonCurrentDeferredLiabilities                       null.Float                                                    `json:"non_current_deferred_liabilities"`
	NonCurrentDeferredRevenue                           null.Float                                                    `json:"non_current_deferred_revenue"`
	NonCurrentDeferredTaxesLiabilities                  null.Float                                                    `json:"non_current_deferred_taxes_liabilities"`
	OtherNonCurrentLiabilities                          null.Float                                                    `json:"other_non_current_liabilities"`
	PreferredSecuritiesOutsideStockEquity               null.Float                                                    `json:"preferred_securities_outside_stock_equity"`
	DerivativeProductLiabilities                        null.Float                                                    `json:"derivative_product_liabilities"`
	CapitalLeaseObligations                             null.Float                                                    `json:"capital_lease_obligations"`
	RestrictedCommonStock                               null.Float                                                    `json:"restricted_common_stock"`
}

// ConsolidatedBalanceSheetLongTermDebtAndCapitalLeaseObligation contains balance sheet data from the consolidated endpoint.
type ConsolidatedBalanceSheetLongTermDebtAndCapitalLeaseObligation struct {
	TotalLongTermDebtAndCapitalLeaseObligation null.Float `json:"total_long_term_debt_and_capital_lease_obligation"`
	LongTermDebt                               null.Float `json:"long_term_debt"`
	LongTermCapitalLeaseObligation             null.Float `json:"long_term_capital_lease_obligation"`
}

// ConsolidatedBalanceSheetEquity contains balance sheet data from the consolidated endpoint.
type ConsolidatedBalanceSheetEquity struct {
	TotalEquityGrossMinorityInterest null.Float                                `json:"total_equity_gross_minority_interest"`
	StockholdersEquity               null.Float                                `json:"stockholders_equity"`
	CommonStockEquity                null.Float                                `json:"common_stock_equity"`
	PreferredStockEquity             null.Float                                `json:"preferred_stock_equity"`
	OtherEquityInterest              null.Float                                `json:"other_equity_interest"`
	MinorityInterest                 null.Float                                `json:"minority_interest"`
	TotalCapitalization              null.Float                                `json:"total_capitalization"`
	NetTangibleAssets                null.Float                                `json:"net_tangible_assets"`
	TangibleBookValue                null.Float                                `json:"tangible_book_value"`
	InvestedCapital                  null.Float                                `json:"invested_capital"`
	WorkingCapital                   null.Float                                `json:"working_capital"`
	CapitalStock                     ConsolidatedBalanceSheetCapitalStock      `json:"capital_stock"`
	EquityAdjustments                ConsolidatedBalanceSheetEquityAdjustments `json:"equity_adjustments"`
	NetDebt                          null.Float                                `json:"net_debt"`
	TotalDebt                        null.Float                                `json:"total_debt"`
}

// ConsolidatedBalanceSheetCapitalStock contains balance sheet data from the consolidated endpoint.
type ConsolidatedBalanceSheetCapitalStock struct {
	CommonStock               null.Float `json:"common_stock"`
	PreferredStock            null.Float `json:"preferred_stock"`
	TotalPartnershipCapital   null.Float `json:"total_partnership_capital"`
	GeneralPartnershipCapital null.Float `json:"general_partnership_capital"`
	LimitedPartnershipCapital null.Float `json:"limited_partnership_capital"`
	CapitalStock              null.Float `json:"capital_stock"`
	OtherCapitalStock         null.Float `json:"other_capital_stock"`
	AdditionalPaidInCapital   null.Float `json:"additional_paid_in_capital"`
	RetainedEarnings          null.Float `json:"retained_earnings"`
	TreasuryStock             null.Float `json:"treasury_stock"`
	TreasurySharesNumber      null.Float `json:"treasury_shares_number"`
	OrdinarySharesNumber      null.Float `json:"ordinary_shares_number"`
	PreferredSharesNumber     null.Float `json:"preferred_shares_number"`
	ShareIssued               null.Float `json:"share_issued"`
}

// ConsolidatedBalanceSheetEquityAdjustments contains balance sheet data from the consolidated endpoint.
type ConsolidatedBalanceSheetEquityAdjustments struct {
	GainsLossesNotAffectingRetainedEarnings null.Float `json:"gains_losses_not_affecting_retained_earnings"`
	OtherEquityAdjustments                  null.Float `json:"other_equity_adjustments"`
	FixedAssetsRevaluationReserve           null.Float `json:"fixed_assets_revaluation_reserve"`
	ForeignCurrencyTranslationAdjustments   null.Float `json:"foreign_currency_translation_adjustments"`
	MinimumPensionLiabilities               null.Float `json:"minimum_pension_liabilities"`
	UnrealizedGainLoss                      null.Float `json:"unrealized_gain_loss"`
}
