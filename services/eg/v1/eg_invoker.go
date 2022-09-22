package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eg/v1/model"
)

type CreateAgenciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAgenciesInvoker) Invoke() (*model.CreateAgenciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAgenciesResponse), nil
	}
}

type CreateChannelInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateChannelInvoker) Invoke() (*model.CreateChannelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateChannelResponse), nil
	}
}

type CreateConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateConnectionInvoker) Invoke() (*model.CreateConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateConnectionResponse), nil
	}
}

type CreateEndpointInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEndpointInvoker) Invoke() (*model.CreateEndpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEndpointResponse), nil
	}
}

type CreateEventSchemaInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEventSchemaInvoker) Invoke() (*model.CreateEventSchemaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEventSchemaResponse), nil
	}
}

type CreateEventSchemaVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEventSchemaVersionInvoker) Invoke() (*model.CreateEventSchemaVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEventSchemaVersionResponse), nil
	}
}

type CreateEventSourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEventSourceInvoker) Invoke() (*model.CreateEventSourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEventSourceResponse), nil
	}
}

type CreateSubscriptionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSubscriptionInvoker) Invoke() (*model.CreateSubscriptionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSubscriptionResponse), nil
	}
}

type CreateSubscriptionTargetInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSubscriptionTargetInvoker) Invoke() (*model.CreateSubscriptionTargetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSubscriptionTargetResponse), nil
	}
}

type DeleteChannelInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteChannelInvoker) Invoke() (*model.DeleteChannelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteChannelResponse), nil
	}
}

type DeleteConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteConnectionInvoker) Invoke() (*model.DeleteConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteConnectionResponse), nil
	}
}

type DeleteEndpointInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEndpointInvoker) Invoke() (*model.DeleteEndpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEndpointResponse), nil
	}
}

type DeleteEventSchemaInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEventSchemaInvoker) Invoke() (*model.DeleteEventSchemaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEventSchemaResponse), nil
	}
}

type DeleteEventSchemaVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEventSchemaVersionInvoker) Invoke() (*model.DeleteEventSchemaVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEventSchemaVersionResponse), nil
	}
}

type DeleteEventSourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEventSourceInvoker) Invoke() (*model.DeleteEventSourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEventSourceResponse), nil
	}
}

type DeleteSubscriptionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSubscriptionInvoker) Invoke() (*model.DeleteSubscriptionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSubscriptionResponse), nil
	}
}

type DeleteSubscriptionTargetInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSubscriptionTargetInvoker) Invoke() (*model.DeleteSubscriptionTargetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSubscriptionTargetResponse), nil
	}
}

type DiscoverEventSchemaFromDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *DiscoverEventSchemaFromDataInvoker) Invoke() (*model.DiscoverEventSchemaFromDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DiscoverEventSchemaFromDataResponse), nil
	}
}

type ListAgenciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAgenciesInvoker) Invoke() (*model.ListAgenciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAgenciesResponse), nil
	}
}

type ListChannelsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListChannelsInvoker) Invoke() (*model.ListChannelsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListChannelsResponse), nil
	}
}

type ListConnectionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConnectionsInvoker) Invoke() (*model.ListConnectionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConnectionsResponse), nil
	}
}

type ListEndpointsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEndpointsInvoker) Invoke() (*model.ListEndpointsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEndpointsResponse), nil
	}
}

type ListEventSchemaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEventSchemaInvoker) Invoke() (*model.ListEventSchemaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEventSchemaResponse), nil
	}
}

type ListEventSchemaVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEventSchemaVersionsInvoker) Invoke() (*model.ListEventSchemaVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEventSchemaVersionsResponse), nil
	}
}

type ListEventSourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEventSourcesInvoker) Invoke() (*model.ListEventSourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEventSourcesResponse), nil
	}
}

type ListEventTargetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEventTargetInvoker) Invoke() (*model.ListEventTargetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEventTargetResponse), nil
	}
}

type ListQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQuotasInvoker) Invoke() (*model.ListQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQuotasResponse), nil
	}
}

type ListSubscriptionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSubscriptionsInvoker) Invoke() (*model.ListSubscriptionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSubscriptionsResponse), nil
	}
}

type ListTriggersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTriggersInvoker) Invoke() (*model.ListTriggersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTriggersResponse), nil
	}
}

type OperateSubscriptionInvoker struct {
	*invoker.BaseInvoker
}

func (i *OperateSubscriptionInvoker) Invoke() (*model.OperateSubscriptionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.OperateSubscriptionResponse), nil
	}
}

type PutEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *PutEventsInvoker) Invoke() (*model.PutEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PutEventsResponse), nil
	}
}

type ShowDetailOfChannelInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDetailOfChannelInvoker) Invoke() (*model.ShowDetailOfChannelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDetailOfChannelResponse), nil
	}
}

type ShowDetailOfConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDetailOfConnectionInvoker) Invoke() (*model.ShowDetailOfConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDetailOfConnectionResponse), nil
	}
}

type ShowDetailOfEventSchemaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDetailOfEventSchemaInvoker) Invoke() (*model.ShowDetailOfEventSchemaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDetailOfEventSchemaResponse), nil
	}
}

type ShowDetailOfEventSchemaVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDetailOfEventSchemaVersionInvoker) Invoke() (*model.ShowDetailOfEventSchemaVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDetailOfEventSchemaVersionResponse), nil
	}
}

type ShowDetailOfEventSourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDetailOfEventSourceInvoker) Invoke() (*model.ShowDetailOfEventSourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDetailOfEventSourceResponse), nil
	}
}

type ShowDetailOfSubscriptionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDetailOfSubscriptionInvoker) Invoke() (*model.ShowDetailOfSubscriptionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDetailOfSubscriptionResponse), nil
	}
}

type ShowDetailOfSubscriptionTargetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDetailOfSubscriptionTargetInvoker) Invoke() (*model.ShowDetailOfSubscriptionTargetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDetailOfSubscriptionTargetResponse), nil
	}
}

type UpdateChannelInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateChannelInvoker) Invoke() (*model.UpdateChannelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateChannelResponse), nil
	}
}

type UpdateConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateConnectionInvoker) Invoke() (*model.UpdateConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateConnectionResponse), nil
	}
}

type UpdateEndpointInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEndpointInvoker) Invoke() (*model.UpdateEndpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEndpointResponse), nil
	}
}

type UpdateEventSchemaInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEventSchemaInvoker) Invoke() (*model.UpdateEventSchemaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEventSchemaResponse), nil
	}
}

type UpdateEventSourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEventSourceInvoker) Invoke() (*model.UpdateEventSourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEventSourceResponse), nil
	}
}

type UpdateSubscriptionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSubscriptionInvoker) Invoke() (*model.UpdateSubscriptionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSubscriptionResponse), nil
	}
}

type UpdateSubscriptionSourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSubscriptionSourceInvoker) Invoke() (*model.UpdateSubscriptionSourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSubscriptionSourceResponse), nil
	}
}

type UpdateSubscriptionTargetInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSubscriptionTargetInvoker) Invoke() (*model.UpdateSubscriptionTargetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSubscriptionTargetResponse), nil
	}
}

type ListApiVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApiVersionsInvoker) Invoke() (*model.ListApiVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApiVersionsResponse), nil
	}
}
