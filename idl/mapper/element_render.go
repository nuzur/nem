package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/element_render"
	pb "github.com/nuzur/nem/idl/gen"
)

func ElementRenderToProto(e main_entity.ElementRender) *pb.ElementRender {
	return &pb.ElementRender{
		X:         e.X,
		Y:         e.Y,
		Width:     e.Width,
		Height:    e.Height,
		Collapsed: e.Collapsed,
	}
}

func ElementRenderSliceToProto(es []main_entity.ElementRender) []*pb.ElementRender {
	res := []*pb.ElementRender{}
	for _, e := range es {
		res = append(res, ElementRenderToProto(e))
	}
	return res
}

func ElementRenderFromProto(m *pb.ElementRender) main_entity.ElementRender {
	if m == nil {
		return main_entity.ElementRender{}
	}
	return main_entity.ElementRender{
		X:         m.X,
		Y:         m.Y,
		Width:     m.Width,
		Height:    m.Height,
		Collapsed: m.GetCollapsed(),
	}
}

func ElementRenderSliceFromProto(es []*pb.ElementRender) []main_entity.ElementRender {
	if es == nil {
		return []main_entity.ElementRender{}
	}
	res := []main_entity.ElementRender{}
	for _, e := range es {
		res = append(res, ElementRenderFromProto(e))
	}
	return res
}
