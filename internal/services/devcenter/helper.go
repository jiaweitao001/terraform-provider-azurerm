package devcenter

import (
	"fmt"
	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/resource-manager/devcenter/2023-04-01/devboxdefinitions"
)

func expandDevBoxDefinitionsSku(input []SkuModel) (*devboxdefinitions.Sku, error) {
	if len(input) == 0 {
		return nil, fmt.Errorf("no sku provided")
	}

	v := input[0]
	output := devboxdefinitions.Sku{
		Capacity: pointer.To(v.Capacity),
		Family:   pointer.To(v.Family),
		Name:     v.Name,
		Size:     pointer.To(v.Size),
		Tier:     pointer.To(devboxdefinitions.SkuTier(v.Tier)),
	}

	return pointer.To(output), nil
}

func expandDevBoxDefinitionsImageReference(input []ImageReferenceModel) (*devboxdefinitions.ImageReference, error) {
	if len(input) == 0 {
		return nil, fmt.Errorf("no image_reference provided")
	}

	v := input[0]
	output := devboxdefinitions.ImageReference{
		Id: pointer.To(v.Id),
	}

	return pointer.To(output), nil
}

func flattenDevBoxDefinitionsSku(input *devboxdefinitions.Sku) []SkuModel {
	if input == nil {
		return nil
	}

	output := SkuModel{
		Capacity: *input.Capacity,
		Family:   *input.Family,
		Name:     input.Name,
		Size:     *input.Size,
		Tier:     string(*input.Tier),
	}

	return []SkuModel{output}
}

func flattenDevBoxDefinitionsImageReference(input *devboxdefinitions.ImageReference) []ImageReferenceModel {
	if input == nil {
		return nil
	}

	output := ImageReferenceModel{
		Id: *input.Id,
	}

	return []ImageReferenceModel{output}
}
