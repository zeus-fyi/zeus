package iris_operators

func (t *IrisOpsTestSuite) TestNestedExtraction() {
	data := map[string]interface{}{
		"jsonrpc": "2.0",
		"result": map[string]interface{}{
			"chain_id": "mainnet",
			"sync_info": map[string]interface{}{
				"latest_block_height": 99656805,
			},
		},
	}

	// Create an instance of the IrisRoutingResponseETL
	r := &IrisRoutingResponseETL{
		ExtractionKey: "result,sync_info,latest_block_height",
	}

	// When: ExtractKeyValue is called
	r.ExtractKeyValue(data)
	// Then: Value and DataType should be set appropriately
	val, ok := ConvertToInt(r.Value)
	t.True(ok)
	t.Equal(99656805, val) // JSON numbers are by default float64 in Go
	t.Equal("int", r.DataType)
}

func (t *IrisOpsTestSuite) TestExtraction() {
	data := map[string]interface{}{
		"jsonrpc": "2.0",
		"result":  99656805,
	}

	// Create an instance of the IrisRoutingResponseETL
	r := &IrisRoutingResponseETL{
		ExtractionKey: "result",
	}

	// When: ExtractKeyValue is called
	r.ExtractKeyValue(data)
	// Then: Value and DataType should be set appropriately
	val, ok := ConvertToInt(r.Value)
	t.True(ok)
	t.Equal(99656805, val) // JSON numbers are by default float64 in Go
	t.Equal("int", r.DataType)
}
