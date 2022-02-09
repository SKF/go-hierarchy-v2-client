package gohierarchyclient

import (
	"context"
	"fmt"
	"net/url"

	"github.com/SKF/go-hierarchy-v2-client/rest/models"

	rest "github.com/SKF/go-rest-utility/client"
	"github.com/SKF/go-utility/v2/stages"
	"github.com/SKF/go-utility/v2/uuid"
)

type TreeFilter struct {
	Depth         int
	Limit         int
	Offset        int
	MetadataKey   string
	MetadataValue string
	NodeTypes     []string
	ModifiedAfter string
}

type HierarchyClient interface {
	GetNode(ctx context.Context, id uuid.UUID) (models.Node, error)
	CreateNode(ctx context.Context, node models.CreateNodeRequest) (models.Node, error)
	UpdateNode(ctx context.Context, id uuid.UUID, node models.UpdateNodeRequest) (models.Node, error)
	DeleteNode(ctx context.Context, id uuid.UUID) error
	DuplicateNode(ctx context.Context, source uuid.UUID, destination uuid.UUID, suffix string) (uuid.UUID, error)

	GetAncestors(ctx context.Context, id uuid.UUID, height int, nodeTypes ...string) ([]models.Node, error)
	GetCompany(ctx context.Context, id uuid.UUID) (models.Node, error)
	GetSubtree(ctx context.Context, id uuid.UUID, filter TreeFilter) ([]models.Node, error)
	GetSubtreeCount(ctx context.Context, id uuid.UUID, nodeTypes ...string) (int64, error)
	getSubtreePage(ctx context.Context, id uuid.UUID, filter TreeFilter, continuationToken string) (models.GetSubtreeResponse, error)

	LockNode(ctx context.Context, id uuid.UUID, recursive bool) error
	UnlockNode(ctx context.Context, id uuid.UUID, recursive bool) error

	GetOrigins(ctx context.Context, provider string) ([]models.Origin, error)
	GetOriginsByType(ctx context.Context, provider, originType string) ([]models.Origin, error)
	GetProviderNodeIDs(ctx context.Context, provider string) ([]uuid.UUID, error)
	GetProviderNodeIDsByType(ctx context.Context, provider, originType string) ([]uuid.UUID, error)
	GetOriginNodeID(ctx context.Context, origin models.Origin) (uuid.UUID, error)
}

type client struct {
	*rest.Client
}

func WithStage(stage string) rest.Option {
	if stage == stages.StageProd {
		return rest.WithBaseURL("https://api.hierarchy.enlight.skf.com/v2/")
	}

	return rest.WithBaseURL(fmt.Sprintf("https://api.%s.hierarchy.enlight.skf.com/v2/", stage))
}

func WithClientID(clientID string) rest.Option {
	clientHeader := "X-Client-Id"
	return rest.WithDefaultHeader(clientHeader, clientID)
}

func NewClient(opts ...rest.Option) HierarchyClient {
	restClient := rest.NewClient(
		append([]rest.Option{
			// Defaults to production stage if no option is supplied
			WithStage(stages.StageProd),
		}, opts...)...,
	)

	return &client{Client: restClient}
}

func (c *client) GetNode(ctx context.Context, id uuid.UUID) (models.Node, error) {
	request := rest.Get("nodes/{node}").
		Assign("node", id).
		SetHeader("Accept", "application/json")

	var response models.GetNodeResponse
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.Node{}, err
	}

	return response.Node, nil
}

func (c *client) CreateNode(ctx context.Context, node models.CreateNodeRequest) (models.Node, error) {
	request := rest.Post("nodes").
		WithJSONPayload(node).
		SetHeader("Accept", "application/json")

	var response models.CreateNodeResponse
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.Node{}, err
	}

	return response.Node, nil
}

func (c *client) UpdateNode(ctx context.Context, id uuid.UUID, node models.UpdateNodeRequest) (models.Node, error) {
	request := rest.Patch("nodes/{node}").
		Assign("node", id).
		WithJSONPayload(node).
		SetHeader("Accept", "application/json")

	var response models.UpdateNodeResponse
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.Node{}, err
	}

	return response.Node, nil
}

func (c *client) DeleteNode(ctx context.Context, id uuid.UUID) (err error) {
	request := rest.Delete("nodes/{node}").
		Assign("node", id).
		SetHeader("Accept", "application/json")

	_, err = c.Do(ctx, request)

	return
}

func (c *client) DuplicateNode(ctx context.Context, source uuid.UUID, destination uuid.UUID, suffix string) (uuid.UUID, error) {
	request := rest.Post("nodes/{node}/duplicate{?destination_id,label_suffix}").
		Assign("node", source).
		Assign("destination_id", destination).
		Assign("label_suffix", suffix).
		SetHeader("Accept", "application/json")

	var response models.DuplicateNodeResponse
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return "", err
	}

	return uuid.UUID(response.NewNodeID), nil
}

func (c *client) GetAncestors(ctx context.Context, id uuid.UUID, height int, nodeTypes ...string) ([]models.Node, error) {
	request := rest.Get("nodes/{node}/ancestors{?height,type*}").
		Assign("node", id).
		Assign("height", height).
		Assign("type", nodeTypes).
		SetHeader("Accept", "application/json")

	var response models.GetAncestorsResponse
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return []models.Node{}, err
	}

	return response.Nodes, nil
}

func (c *client) GetCompany(ctx context.Context, id uuid.UUID) (models.Node, error) {
	request := rest.Get("nodes/{node}/company").
		Assign("node", id).
		SetHeader("Accept", "application/json")

	var response models.GetCompanyResponse
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.Node{}, err
	}

	return response.Node, nil
}

func (c *client) getSubtreePage(ctx context.Context, id uuid.UUID, filter TreeFilter, continuationToken string) (models.GetSubtreeResponse, error) {
	request := rest.Get("nodes/{node}/subtree{?depth,limit,offset,metadata_key,metadata_value,continuation_token,type*}").
		Assign("node", id).
		Assign("depth", filter.Depth).
		Assign("limit", filter.Limit).
		Assign("offset", filter.Offset).
		Assign("metadata_key", filter.MetadataKey).
		Assign("metadata_value", filter.MetadataValue).
		Assign("type", filter.NodeTypes).
		Assign("continuation_token", continuationToken).
		Assign("modified_after", filter.ModifiedAfter).
		SetHeader("Accept", "application/json")

	var response models.GetSubtreeResponse
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.GetSubtreeResponse{}, err
	}

	return response, nil
}

func (c *client) GetSubtree(ctx context.Context, id uuid.UUID, filter TreeFilter) ([]models.Node, error) {
	response, err := c.getSubtreePage(ctx, id, filter, "")
	if err != nil {
		return []models.Node{}, err
	}

	nodes := response.Nodes

	for response.Links != nil && response.Links.Next != "" {
		continuationToken, err := getContinuationToken(response.Links.Next)
		if err != nil {
			return []models.Node{}, err
		}

		response, err = c.getSubtreePage(ctx, id, filter, continuationToken)
		if err != nil {
			return []models.Node{}, err
		}

		nodes = append(nodes, response.Nodes...)
	}

	return nodes, nil
}

func (c *client) GetSubtreeCount(ctx context.Context, id uuid.UUID, nodeTypes ...string) (int64, error) {
	request := rest.Get("nodes/{node}/subtree/count{?type*}").
		Assign("node", id).
		Assign("type", nodeTypes).
		SetHeader("Accept", "application/json")

	var response models.GetSubtreeCountResponse
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return 0, err
	}

	return response.Count, nil
}

func (c *client) getOriginsPage(ctx context.Context, provider, continuationToken string, limit int) (models.GetOriginsResponse, error) {
	request := rest.Get("origins/{provider,continuation_token,limit}").
		Assign("provider", provider).
		Assign("continuation_token", continuationToken).
		Assign("limit", limit).
		SetHeader("Accept", "application/json")

	var response models.GetOriginsResponse
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.GetOriginsResponse{}, err
	}

	return response, nil
}

func (c *client) GetOrigins(ctx context.Context, provider string) ([]models.Origin, error) {
	response, err := c.getOriginsPage(ctx, provider, "", 0)
	if err != nil {
		return []models.Origin{}, err
	}

	originsPtrs := response.Origins

	for response.Links != nil && response.Links.Next != "" {
		continuationToken, err := getContinuationToken(response.Links.Next)
		if err != nil {
			return []models.Origin{}, err
		}

		response, err = c.getOriginsPage(ctx, provider, continuationToken, 0)
		if err != nil {
			return []models.Origin{}, err
		}

		originsPtrs = append(originsPtrs, response.Origins...)
	}

	origins := make([]models.Origin, 0, len(originsPtrs))

	for _, o := range originsPtrs {
		if o != nil {
			origins = append(origins, *o)
		}
	}

	return origins, nil
}

func (c *client) getOriginsByTypePage(ctx context.Context, provider, originType, continuationToken string, limit int) (models.GetOriginsResponse, error) {
	request := rest.Get("origins/{provider}/{type}").
		Assign("provider", provider).
		Assign("type", originType).
		Assign("continuation_token", continuationToken).
		Assign("limit", limit).
		SetHeader("Accept", "application/json")

	var response models.GetOriginsResponse
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.GetOriginsResponse{}, err
	}

	return response, nil
}

func (c *client) GetOriginsByType(ctx context.Context, provider, originType string) ([]models.Origin, error) {
	response, err := c.getOriginsByTypePage(ctx, provider, originType, "", 0)
	if err != nil {
		return []models.Origin{}, err
	}

	originsPtrs := response.Origins

	for response.Links != nil && response.Links.Next != "" {
		continuationToken, err := getContinuationToken(response.Links.Next)
		if err != nil {
			return []models.Origin{}, err
		}

		response, err = c.getOriginsByTypePage(ctx, provider, originType, continuationToken, 0)
		if err != nil {
			return []models.Origin{}, err
		}

		originsPtrs = append(originsPtrs, response.Origins...)
	}

	origins := make([]models.Origin, 0, len(originsPtrs))

	for _, o := range originsPtrs {
		if o != nil {
			origins = append(origins, *o)
		}
	}

	return origins, nil
}

func (c *client) getProviderNodeIDsPage(ctx context.Context, provider, continuationToken string, limit int) (models.GetNodesByPartialOriginResponse, error) {
	request := rest.Get("origins/{provider}/nodes").
		Assign("provider", provider).
		Assign("continuation_token", continuationToken).
		Assign("limit", limit).
		SetHeader("Accept", "application/json")

	var response models.GetNodesByPartialOriginResponse
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.GetNodesByPartialOriginResponse{}, err
	}

	return response, nil
}

func (c *client) GetProviderNodeIDs(ctx context.Context, provider string) ([]uuid.UUID, error) {
	response, err := c.getProviderNodeIDsPage(ctx, provider, "", 0)
	if err != nil {
		return []uuid.UUID{}, err
	}

	nodeIDs := []uuid.UUID{}
	for _, nID := range response.NodeIds {
		nodeIDs = append(nodeIDs, uuid.UUID(nID))
	}

	for response.Links != nil && response.Links.Next != "" {
		continuationToken, err := getContinuationToken(response.Links.Next)
		if err != nil {
			return []uuid.UUID{}, err
		}

		response, err = c.getProviderNodeIDsPage(ctx, provider, continuationToken, 0)
		if err != nil {
			return []uuid.UUID{}, err
		}

		for _, nID := range response.NodeIds {
			nodeIDs = append(nodeIDs, uuid.UUID(nID))
		}
	}

	return nodeIDs, nil
}

func (c *client) getProviderNodeIDsByTypePage(ctx context.Context, provider, originType, continuationToken string, limit int) (models.GetNodesByPartialOriginResponse, error) {
	request := rest.Get("origins/{provider}/{type}/nodes").
		Assign("provider", provider).
		Assign("type", originType).
		Assign("continuation_token", continuationToken).
		Assign("limit", limit).
		SetHeader("Accept", "application/json")

	var response models.GetNodesByPartialOriginResponse
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.GetNodesByPartialOriginResponse{}, err
	}

	return response, nil
}

func (c *client) GetProviderNodeIDsByType(ctx context.Context, provider, originType string) ([]uuid.UUID, error) {
	response, err := c.getProviderNodeIDsByTypePage(ctx, provider, originType, "", 0)
	if err != nil {
		return []uuid.UUID{}, err
	}

	nodeIDs := []uuid.UUID{}
	for _, nID := range response.NodeIds {
		nodeIDs = append(nodeIDs, uuid.UUID(nID))
	}

	for response.Links != nil && response.Links.Next != "" {
		continuationToken, err := getContinuationToken(response.Links.Next)
		if err != nil {
			return []uuid.UUID{}, err
		}

		response, err = c.getProviderNodeIDsByTypePage(ctx, provider, originType, continuationToken, 0)
		if err != nil {
			return []uuid.UUID{}, err
		}

		for _, nID := range response.NodeIds {
			nodeIDs = append(nodeIDs, uuid.UUID(nID))
		}
	}

	return nodeIDs, nil
}

func (c *client) GetOriginNodeID(ctx context.Context, origin models.Origin) (uuid.UUID, error) {
	if err := validateOriginInput(origin); err != nil {
		return "", err
	}

	request := rest.Get("origins/{provider}/{type}/{id}/node").
		Assign("provider", *origin.Provider).
		Assign("type", *origin.Type).
		Assign("id", *origin.ID).
		SetHeader("Accept", "application/json")

	var response models.GetNodeByOriginResponse
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return "", err
	}

	return response.NodeID, nil
}

func (c *client) LockNode(ctx context.Context, id uuid.UUID, recursive bool) error {
	request := rest.Put("nodes/{node}/lock{?recursive}").
		Assign("node", id).
		Assign("recursive", recursive).
		SetHeader("Accept", "application/json")

	_, err := c.Do(ctx, request)

	return err
}
func (c *client) UnlockNode(ctx context.Context, id uuid.UUID, recursive bool) error {
	request := rest.Delete("nodes/{node}/lock{?recursive}").
		Assign("node", id).
		Assign("recursive", recursive).
		SetHeader("Accept", "application/json")

	_, err := c.Do(ctx, request)

	return err
}

func getContinuationToken(nextLink string) (string, error) {
	queryValues, err := url.ParseQuery(nextLink)
	if err != nil {
		err = fmt.Errorf("unable to parse next Link:'%s' as an URL %w", nextLink, err)
		return "", err
	}

	continuationToken := queryValues.Get("continuation_token")
	if continuationToken == "" {
		err = fmt.Errorf("expected query value continuation_token not found in next Link: %s", nextLink)
		return "", err
	}

	return continuationToken, nil
}

func validateOriginInput(origin models.Origin) error {
	if origin.Type == nil || origin.Provider == nil || origin.ID == nil {
		return fmt.Errorf("invalid origin one or more fields is nil")
	}

	return nil
}
