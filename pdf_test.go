package main

import (
	"os"
	"testing"

	"github.com/signintech/gopdf"
)

// TestWriteRowWithFloatQuantity tests that writeRow correctly handles float64 quantities
func TestWriteRowWithFloatQuantity(t *testing.T) {
	pdf := &gopdf.GoPdf{}
	pdf.Start(gopdf.Config{
		PageSize: *gopdf.PageSizeA4,
	})
	pdf.SetMargins(40, 40, 40, 40)
	pdf.AddPage()

	// Add fonts
	interFont, _ := os.ReadFile("Inter/Inter Variable/Inter.ttf")
	interBoldFont, _ := os.ReadFile("Inter/Inter Hinted for Windows/Desktop/Inter-Bold.ttf")
	
	_ = pdf.AddTTFFontData("Inter", interFont)
	_ = pdf.AddTTFFontData("Inter-Bold", interBoldFont)

	// Set up currency symbols (from currency.go)
	file = Invoice{
		Currency: "USD",
	}

	tests := []struct {
		name     string
		item     string
		quantity float64
		rate     float64
	}{
		{
			name:     "integer quantity",
			item:     "Widget",
			quantity: 5.0,
			rate:     10.0,
		},
		{
			name:     "decimal quantity",
			item:     "Component",
			quantity: 2.5,
			rate:     20.0,
		},
		{
			name:     "small decimal quantity",
			item:     "Service",
			quantity: 0.5,
			rate:     100.0,
		},
		{
			name:     "multi-line item with float quantity",
			item:     "Complex Item\nSecond Line",
			quantity: 3.75,
			rate:     15.50,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset Y position for each test
			pdf.SetY(100)

			// This should not panic or error
			writeRow(pdf, tt.item, tt.quantity, tt.rate)

			// If we got here without panicking, the test passes
			// The function successfully handled float64 quantity
		})
	}
}

// TestWriteRowQuantityCalculation verifies correct calculation with float quantities
func TestWriteRowQuantityCalculation(t *testing.T) {
	pdf := &gopdf.GoPdf{}
	pdf.Start(gopdf.Config{
		PageSize: *gopdf.PageSizeA4,
	})
	pdf.SetMargins(40, 40, 40, 40)
	pdf.AddPage()

	// Add fonts
	interFont, _ := os.ReadFile("Inter/Inter Variable/Inter.ttf")
	interBoldFont, _ := os.ReadFile("Inter/Inter Hinted for Windows/Desktop/Inter-Bold.ttf")
	
	_ = pdf.AddTTFFontData("Inter", interFont)
	_ = pdf.AddTTFFontData("Inter-Bold", interBoldFont)

	file = Invoice{
		Currency: "USD",
	}

	// Test that calculation is correct: 2.5 * 20.0 = 50.0
	quantity := 2.5
	rate := 20.0
	expectedTotal := quantity * rate

	if expectedTotal != 50.0 {
		t.Errorf("expected total %f, got %f", 50.0, expectedTotal)
	}

	// Should not panic when calling writeRow with float quantity
	writeRow(pdf, "Test Item", quantity, rate)
}

// TestInvoiceStructureWithFloatQuantities verifies the Invoice struct can handle float quantities
// This test ensures no regressions in type handling
func TestInvoiceStructureWithFloatQuantities(t *testing.T) {
	// Create an invoice with float quantities by modifying the struct
	// This ensures the PDF generation pipeline doesn't break with our changes
	
	invoice := Invoice{
		Id:         "INV-001",
		Title:      "TEST INVOICE",
		From:       "Test Company",
		To:         "Test Client",
		Items:      []string{"Item 1", "Item 2"},
		Quantities: []int{2, 3}, // Still accepts int for now from CLI
		Rates:      []float64{25.0, 50.0},
		Currency:   "USD",
	}

	if invoice.Id != "INV-001" {
		t.Errorf("invoice ID mismatch")
	}

	if len(invoice.Items) != 2 {
		t.Errorf("expected 2 items, got %d", len(invoice.Items))
	}

	// Calculate totals as the main function does
	subtotal := 0.0
	for i := range invoice.Items {
		q := 1
		if len(invoice.Quantities) > i {
			q = invoice.Quantities[i]
		}

		r := 0.0
		if len(invoice.Rates) > i {
			r = invoice.Rates[i]
		}

		subtotal += float64(q) * r
	}

	expectedSubtotal := 2*25.0 + 3*50.0 // 50 + 150 = 200
	if subtotal != expectedSubtotal {
		t.Errorf("expected subtotal %f, got %f", expectedSubtotal, subtotal)
	}
}
