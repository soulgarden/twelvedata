package response_test

import (
	"strings"
	"testing"

	"github.com/soulgarden/twelvedata/response"
)

func TestDocumentedFractionalNumbers(t *testing.T) {
	tests := []struct {
		name, data string
		newTarget  func() any
	}{
		{"/splits/splits.from_factor", `{"splits":[{"from_factor":1.5}]}`, func() any { return new(response.Splits) }},
		{"/splits/splits.to_factor", `{"splits":[{"to_factor":1.5}]}`, func() any { return new(response.Splits) }},
		{"/splits_calendar/from_factor", `[{"from_factor":1.5}]`, func() any { return new(response.SplitsCalendar) }},
		{"/splits_calendar/to_factor", `[{"to_factor":1.5}]`, func() any { return new(response.SplitsCalendar) }},
		{"/statistics/statistics.valuations_metrics.market_capitalization", `{"statistics":{"valuations_metrics":{"market_capitalization":1.5}}}`, func() any { return new(response.Statistics) }},
		{"/statistics/statistics.valuations_metrics.enterprise_value", `{"statistics":{"valuations_metrics":{"enterprise_value":1.5}}}`, func() any { return new(response.Statistics) }},
		{"/statistics/statistics.financials.income_statement.revenue_ttm", `{"statistics":{"financials":{"income_statement":{"revenue_ttm":1.5}}}}`, func() any { return new(response.Statistics) }},
		{"/statistics/statistics.financials.income_statement.gross_profit_ttm", `{"statistics":{"financials":{"income_statement":{"gross_profit_ttm":1.5}}}}`, func() any { return new(response.Statistics) }},
		{"/statistics/statistics.financials.income_statement.ebitda", `{"statistics":{"financials":{"income_statement":{"ebitda":1.5}}}}`, func() any { return new(response.Statistics) }},
		{"/statistics/statistics.financials.income_statement.net_income_to_common_ttm", `{"statistics":{"financials":{"income_statement":{"net_income_to_common_ttm":1.5}}}}`, func() any { return new(response.Statistics) }},
		{"/statistics/statistics.financials.balance_sheet.total_cash_mrq", `{"statistics":{"financials":{"balance_sheet":{"total_cash_mrq":1.5}}}}`, func() any { return new(response.Statistics) }},
		{"/statistics/statistics.financials.balance_sheet.total_debt_mrq", `{"statistics":{"financials":{"balance_sheet":{"total_debt_mrq":1.5}}}}`, func() any { return new(response.Statistics) }},
		{"/statistics/statistics.financials.cash_flow.operating_cash_flow_ttm", `{"statistics":{"financials":{"cash_flow":{"operating_cash_flow_ttm":1.5}}}}`, func() any { return new(response.Statistics) }},
		{"/statistics/statistics.financials.cash_flow.levered_free_cash_flow_ttm", `{"statistics":{"financials":{"cash_flow":{"levered_free_cash_flow_ttm":1.5}}}}`, func() any { return new(response.Statistics) }},
		{"/statistics/statistics.stock_statistics.shares_outstanding", `{"statistics":{"stock_statistics":{"shares_outstanding":1.5}}}`, func() any { return new(response.Statistics) }},
		{"/statistics/statistics.stock_statistics.float_shares", `{"statistics":{"stock_statistics":{"float_shares":1.5}}}`, func() any { return new(response.Statistics) }},
		{"/balance_sheet/balance_sheet.assets.current_assets.cash", `{"balance_sheet":[{"assets":{"current_assets":{"cash":1.5}}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.assets.current_assets.cash_equivalents", `{"balance_sheet":[{"assets":{"current_assets":{"cash_equivalents":1.5}}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.assets.current_assets.cash_and_cash_equivalents", `{"balance_sheet":[{"assets":{"current_assets":{"cash_and_cash_equivalents":1.5}}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.assets.current_assets.other_short_term_investments", `{"balance_sheet":[{"assets":{"current_assets":{"other_short_term_investments":1.5}}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.assets.current_assets.accounts_receivable", `{"balance_sheet":[{"assets":{"current_assets":{"accounts_receivable":1.5}}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.assets.current_assets.other_receivables", `{"balance_sheet":[{"assets":{"current_assets":{"other_receivables":1.5}}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.assets.current_assets.inventory", `{"balance_sheet":[{"assets":{"current_assets":{"inventory":1.5}}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.assets.current_assets.prepaid_assets", `{"balance_sheet":[{"assets":{"current_assets":{"prepaid_assets":1.5}}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.assets.current_assets.restricted_cash", `{"balance_sheet":[{"assets":{"current_assets":{"restricted_cash":1.5}}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.assets.current_assets.assets_held_for_sale", `{"balance_sheet":[{"assets":{"current_assets":{"assets_held_for_sale":1.5}}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.assets.current_assets.hedging_assets", `{"balance_sheet":[{"assets":{"current_assets":{"hedging_assets":1.5}}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.assets.current_assets.other_current_assets", `{"balance_sheet":[{"assets":{"current_assets":{"other_current_assets":1.5}}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.assets.current_assets.total_current_assets", `{"balance_sheet":[{"assets":{"current_assets":{"total_current_assets":1.5}}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.assets.non_current_assets.properties", `{"balance_sheet":[{"assets":{"non_current_assets":{"properties":1.5}}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.assets.non_current_assets.land_and_improvements", `{"balance_sheet":[{"assets":{"non_current_assets":{"land_and_improvements":1.5}}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.assets.non_current_assets.machinery_furniture_equipment", `{"balance_sheet":[{"assets":{"non_current_assets":{"machinery_furniture_equipment":1.5}}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.assets.non_current_assets.construction_in_progress", `{"balance_sheet":[{"assets":{"non_current_assets":{"construction_in_progress":1.5}}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.assets.non_current_assets.leases", `{"balance_sheet":[{"assets":{"non_current_assets":{"leases":1.5}}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.assets.non_current_assets.accumulated_depreciation", `{"balance_sheet":[{"assets":{"non_current_assets":{"accumulated_depreciation":1.5}}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.assets.non_current_assets.goodwill", `{"balance_sheet":[{"assets":{"non_current_assets":{"goodwill":1.5}}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.assets.non_current_assets.investment_properties", `{"balance_sheet":[{"assets":{"non_current_assets":{"investment_properties":1.5}}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.assets.non_current_assets.financial_assets", `{"balance_sheet":[{"assets":{"non_current_assets":{"financial_assets":1.5}}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.assets.non_current_assets.intangible_assets", `{"balance_sheet":[{"assets":{"non_current_assets":{"intangible_assets":1.5}}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.assets.non_current_assets.investments_and_advances", `{"balance_sheet":[{"assets":{"non_current_assets":{"investments_and_advances":1.5}}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.assets.non_current_assets.other_non_current_assets", `{"balance_sheet":[{"assets":{"non_current_assets":{"other_non_current_assets":1.5}}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.assets.non_current_assets.total_non_current_assets", `{"balance_sheet":[{"assets":{"non_current_assets":{"total_non_current_assets":1.5}}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.assets.total_assets", `{"balance_sheet":[{"assets":{"total_assets":1.5}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.liabilities.current_liabilities.accounts_payable", `{"balance_sheet":[{"liabilities":{"current_liabilities":{"accounts_payable":1.5}}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.liabilities.current_liabilities.accrued_expenses", `{"balance_sheet":[{"liabilities":{"current_liabilities":{"accrued_expenses":1.5}}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.liabilities.current_liabilities.short_term_debt", `{"balance_sheet":[{"liabilities":{"current_liabilities":{"short_term_debt":1.5}}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.liabilities.current_liabilities.deferred_revenue", `{"balance_sheet":[{"liabilities":{"current_liabilities":{"deferred_revenue":1.5}}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.liabilities.current_liabilities.tax_payable", `{"balance_sheet":[{"liabilities":{"current_liabilities":{"tax_payable":1.5}}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.liabilities.current_liabilities.pensions", `{"balance_sheet":[{"liabilities":{"current_liabilities":{"pensions":1.5}}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.liabilities.current_liabilities.other_current_liabilities", `{"balance_sheet":[{"liabilities":{"current_liabilities":{"other_current_liabilities":1.5}}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.liabilities.current_liabilities.total_current_liabilities", `{"balance_sheet":[{"liabilities":{"current_liabilities":{"total_current_liabilities":1.5}}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.liabilities.non_current_liabilities.long_term_provisions", `{"balance_sheet":[{"liabilities":{"non_current_liabilities":{"long_term_provisions":1.5}}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.liabilities.non_current_liabilities.long_term_debt", `{"balance_sheet":[{"liabilities":{"non_current_liabilities":{"long_term_debt":1.5}}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.liabilities.non_current_liabilities.provision_for_risks_and_charges", `{"balance_sheet":[{"liabilities":{"non_current_liabilities":{"provision_for_risks_and_charges":1.5}}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.liabilities.non_current_liabilities.deferred_liabilities", `{"balance_sheet":[{"liabilities":{"non_current_liabilities":{"deferred_liabilities":1.5}}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.liabilities.non_current_liabilities.derivative_product_liabilities", `{"balance_sheet":[{"liabilities":{"non_current_liabilities":{"derivative_product_liabilities":1.5}}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.liabilities.non_current_liabilities.other_non_current_liabilities", `{"balance_sheet":[{"liabilities":{"non_current_liabilities":{"other_non_current_liabilities":1.5}}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.liabilities.non_current_liabilities.total_non_current_liabilities", `{"balance_sheet":[{"liabilities":{"non_current_liabilities":{"total_non_current_liabilities":1.5}}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.liabilities.total_liabilities", `{"balance_sheet":[{"liabilities":{"total_liabilities":1.5}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.shareholders_equity.common_stock", `{"balance_sheet":[{"shareholders_equity":{"common_stock":1.5}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.shareholders_equity.retained_earnings", `{"balance_sheet":[{"shareholders_equity":{"retained_earnings":1.5}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.shareholders_equity.other_shareholders_equity", `{"balance_sheet":[{"shareholders_equity":{"other_shareholders_equity":1.5}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.shareholders_equity.total_shareholders_equity", `{"balance_sheet":[{"shareholders_equity":{"total_shareholders_equity":1.5}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.shareholders_equity.additional_paid_in_capital", `{"balance_sheet":[{"shareholders_equity":{"additional_paid_in_capital":1.5}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.shareholders_equity.treasury_stock", `{"balance_sheet":[{"shareholders_equity":{"treasury_stock":1.5}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/balance_sheet/balance_sheet.shareholders_equity.minority_interest", `{"balance_sheet":[{"shareholders_equity":{"minority_interest":1.5}}]}`, func() any { return new(response.BalanceSheets) }},
		{"/cash_flow/cash_flow.operating_activities.net_income", `{"cash_flow":[{"operating_activities":{"net_income":1.5}}]}`, func() any { return new(response.CashFlows) }},
		{"/cash_flow/cash_flow.operating_activities.depreciation", `{"cash_flow":[{"operating_activities":{"depreciation":1.5}}]}`, func() any { return new(response.CashFlows) }},
		{"/cash_flow/cash_flow.operating_activities.deferred_taxes", `{"cash_flow":[{"operating_activities":{"deferred_taxes":1.5}}]}`, func() any { return new(response.CashFlows) }},
		{"/cash_flow/cash_flow.operating_activities.stock_based_compensation", `{"cash_flow":[{"operating_activities":{"stock_based_compensation":1.5}}]}`, func() any { return new(response.CashFlows) }},
		{"/cash_flow/cash_flow.operating_activities.other_non_cash_items", `{"cash_flow":[{"operating_activities":{"other_non_cash_items":1.5}}]}`, func() any { return new(response.CashFlows) }},
		{"/cash_flow/cash_flow.operating_activities.accounts_receivable", `{"cash_flow":[{"operating_activities":{"accounts_receivable":1.5}}]}`, func() any { return new(response.CashFlows) }},
		{"/cash_flow/cash_flow.operating_activities.accounts_payable", `{"cash_flow":[{"operating_activities":{"accounts_payable":1.5}}]}`, func() any { return new(response.CashFlows) }},
		{"/cash_flow/cash_flow.operating_activities.other_assets_liabilities", `{"cash_flow":[{"operating_activities":{"other_assets_liabilities":1.5}}]}`, func() any { return new(response.CashFlows) }},
		{"/cash_flow/cash_flow.operating_activities.operating_cash_flow", `{"cash_flow":[{"operating_activities":{"operating_cash_flow":1.5}}]}`, func() any { return new(response.CashFlows) }},
		{"/cash_flow/cash_flow.investing_activities.capital_expenditures", `{"cash_flow":[{"investing_activities":{"capital_expenditures":1.5}}]}`, func() any { return new(response.CashFlows) }},
		{"/cash_flow/cash_flow.investing_activities.net_intangibles", `{"cash_flow":[{"investing_activities":{"net_intangibles":1.5}}]}`, func() any { return new(response.CashFlows) }},
		{"/cash_flow/cash_flow.investing_activities.net_acquisitions", `{"cash_flow":[{"investing_activities":{"net_acquisitions":1.5}}]}`, func() any { return new(response.CashFlows) }},
		{"/cash_flow/cash_flow.investing_activities.purchase_of_investments", `{"cash_flow":[{"investing_activities":{"purchase_of_investments":1.5}}]}`, func() any { return new(response.CashFlows) }},
		{"/cash_flow/cash_flow.investing_activities.sale_of_investments", `{"cash_flow":[{"investing_activities":{"sale_of_investments":1.5}}]}`, func() any { return new(response.CashFlows) }},
		{"/cash_flow/cash_flow.investing_activities.other_investing_activity", `{"cash_flow":[{"investing_activities":{"other_investing_activity":1.5}}]}`, func() any { return new(response.CashFlows) }},
		{"/cash_flow/cash_flow.investing_activities.investing_cash_flow", `{"cash_flow":[{"investing_activities":{"investing_cash_flow":1.5}}]}`, func() any { return new(response.CashFlows) }},
		{"/cash_flow/cash_flow.financing_activities.long_term_debt_issuance", `{"cash_flow":[{"financing_activities":{"long_term_debt_issuance":1.5}}]}`, func() any { return new(response.CashFlows) }},
		{"/cash_flow/cash_flow.financing_activities.long_term_debt_payments", `{"cash_flow":[{"financing_activities":{"long_term_debt_payments":1.5}}]}`, func() any { return new(response.CashFlows) }},
		{"/cash_flow/cash_flow.financing_activities.short_term_debt_issuance", `{"cash_flow":[{"financing_activities":{"short_term_debt_issuance":1.5}}]}`, func() any { return new(response.CashFlows) }},
		{"/cash_flow/cash_flow.financing_activities.common_stock_issuance", `{"cash_flow":[{"financing_activities":{"common_stock_issuance":1.5}}]}`, func() any { return new(response.CashFlows) }},
		{"/cash_flow/cash_flow.financing_activities.common_stock_repurchase", `{"cash_flow":[{"financing_activities":{"common_stock_repurchase":1.5}}]}`, func() any { return new(response.CashFlows) }},
		{"/cash_flow/cash_flow.financing_activities.common_dividends", `{"cash_flow":[{"financing_activities":{"common_dividends":1.5}}]}`, func() any { return new(response.CashFlows) }},
		{"/cash_flow/cash_flow.financing_activities.other_financing_charges", `{"cash_flow":[{"financing_activities":{"other_financing_charges":1.5}}]}`, func() any { return new(response.CashFlows) }},
		{"/cash_flow/cash_flow.financing_activities.financing_cash_flow", `{"cash_flow":[{"financing_activities":{"financing_cash_flow":1.5}}]}`, func() any { return new(response.CashFlows) }},
		{"/cash_flow/cash_flow.end_cash_position", `{"cash_flow":[{"end_cash_position":1.5}]}`, func() any { return new(response.CashFlows) }},
		{"/cash_flow/cash_flow.income_tax_paid", `{"cash_flow":[{"income_tax_paid":1.5}]}`, func() any { return new(response.CashFlows) }},
		{"/cash_flow/cash_flow.interest_paid", `{"cash_flow":[{"interest_paid":1.5}]}`, func() any { return new(response.CashFlows) }},
		{"/cash_flow/cash_flow.free_cash_flow", `{"cash_flow":[{"free_cash_flow":1.5}]}`, func() any { return new(response.CashFlows) }},
		{"/bbands/meta.indicator.sd", `{"meta":{"indicator":{"sd":1.5}}}`, func() any { return new(response.BBands) }},
		{"/percent_b/meta.indicator.sd", `{"meta":{"indicator":{"sd":1.5}}}`, func() any { return new(response.PercentB) }},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, number := range []string{"1.5", `"1.5"`, "0", "null"} {
				t.Run(number, func(t *testing.T) {
					assertContractRoundTrip(t, []byte(strings.ReplaceAll(tt.data, "1.5", number)), tt.newTarget())
				})
			}
		})
	}
}
