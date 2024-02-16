---
sidebar_position: 8
displayed_sidebar: mockingbird
---

# Dynamic Context Sizing

## Introduction

This documentation provides detailed insights into the prompt context management options. It specifically addresses the
handling of token overflow for different AI models, covering system capabilities, user configurations, and expected
behaviors.

## Contents

1. [Model Limit Checks](#model-limit-checks)
2. [Input Breakdown and Re-chunking](#input-breakdown-and-re-chunking)
3. [Expected Behavior for Users](#expected-behavior-for-users)
4. [Truncate vs Deduce](#truncate-vs-deduce)
5. [Margin Buffer Sizing](#margin-buffer-sizing)
6. [Flow Diagram](#flow-diagram)

## Model Limit Checks

The service conducts thorough checks of each input against the predefined token limits for various models, ensuring
compatibility and preventing processing errors. Supported models include various versions of GPT-3.5 and GPT-4, each
with specific token limits.

## Input Breakdown and Re-chunking

When input exceeds model limits, the system employs two strategies to manage overflow:

- **Truncate:** Reduces input to fit within limits, potentially losing information.
- **Deduce:** Breaks down input into manageable chunks, preserving as much information as possible.

## Expected Behavior for Users

Users can configure the token overflow management system by setting the margin buffer size and selecting between
truncation and deduction. This flexibility allows for tailored behavior, depending on the need for precision or
information preservation.

If the results are from search indexer retrievals, or json schema formatted, then the results can be expected to
preserve their original formatting.

If the results are in string format, then the results are not guaranteed to preserve their original formatting when
re-chunked, so if
the result was a long string of a json object array, it may result in a string output chunk that is not valid json.

## Truncate vs Deduce

- **Truncate:** A direct approach where input exceeding the token limit is cut off.
- **Deduce:** An intelligent breakdown of input into smaller chunks, each conforming to token limits.

## Margin Buffer Sizing

The margin buffer is a critical user-configurable parameter that defines a threshold for action on inputs near the token
limit. It's expressed as a percentage, providing a safety margin to accommodate token count estimation uncertainties.

- **Minimum Margin Buffer:** 0.2 (20%)
- **Maximum Margin Buffer:** 0.8 (80%)

Setting to 0.5 (50%) means that you are reserving 50% of the input context size for your prompt body input, so if it
exceeds this it will re-chunk the input into smaller pieces.

## Flow Diagram

The following Mermaid diagram illustrates the decision process between truncating and deducing input based on token
overflow and margin buffer settings:

```mermaid
flowchart TD
    A[Start] --> B{Input Exceeds Model Limit?}
    B -- Yes --> C{Margin Buffer Check}
    C -- Within Range --> D{Token Overflow Strategy}
    D -- Deduce --> E[Break Input Into Chunks]
    D -- Truncate --> F[Cut Off Excess Input]
    C -- Out of Range --> G[Adjust Margin Buffer]
    B -- No --> H[Proceed With Input]
    E --> I[End]
    F --> I
    G --> D
    H --> I
