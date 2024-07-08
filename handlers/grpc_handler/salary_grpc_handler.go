package grpchandler

import (
	"context"
	"fmt"
	"salary_project/entity"
	"salary_project/entity/proto"
	"salary_project/service"
)

type SalaryGrpcHandler struct {
	proto.UnimplementedEmployeeToSalaryServer
	SalaryService *service.SalaryService
}

func (s *SalaryGrpcHandler) CreateSalary(ctx context.Context, in *proto.CreateSalaryRequest) (*proto.CreateSalaryResponse, error) {

	fmt.Println(in)
	//fmt.Println("DDDDDDDDDDDDDDDDDDDDDDDDD")
	payload := entity.CreateEmployeeSalary{
		Salary_Amount: int(in.SalaryAmount),
		Joining_Date:  in.JoiningDate,
		Project:       in.Project,
		Employee_Id:   in.EmployeeId,
	}
	res, err := s.SalaryService.AddSalaryService(ctx, payload)
	if err != nil {
		fmt.Println(err)
		return &proto.CreateSalaryResponse{}, err
	}

	fmt.Println(res)

	resp := &proto.CreateSalaryResponse{
		Status:  true,
		Message: "Salary Created successfully",
	}

	return resp, nil
}
